package api

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/GotoRen/storage-api-sample/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Client struct {
	S3 *s3.S3
}

// Upload uploads any file to remote bucket.
func (client *Client) Upload(cfg *config.Config) error {
	file, err := os.Open(cfg.UploadFilePath)
	if err != nil {
		return fmt.Errorf("failed to get local file: %w:", err)
	}
	defer file.Close()

	uploader := s3manager.NewUploaderWithClient(client.S3)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(cfg.ObjectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload local file: %w:", err)
	}

	log.Println("[DEBUG] Upload successfully:", cfg.UploadFilePath)

	return nil
}

// Read reads any file from remote bucket.
func (client *Client) Read(cfg *config.Config) error {
	objOutput, err := client.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(cfg.ObjectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to get object information: %w", err)
	}
	defer objOutput.Body.Close()

	bodyBytes, err := io.ReadAll(objOutput.Body)
	if err != nil {
		return fmt.Errorf("failed to read object body: %w", err)
	}

	contents := string(bodyBytes)

	log.Print("[DEBUG] Object body: ", contents)

	return nil
}

// Download downloads any file from remote bucket.
func (client *Client) Download(cfg *config.Config) error {
	var token *string
	for complete := false; !complete; {
		input := s3.ListObjectsV2Input{
			Bucket:            &cfg.BucketName,
			ContinuationToken: token,
		}

		output, err := client.S3.ListObjectsV2(&input)
		if err != nil {
			return fmt.Errorf("failed to get bucket information: %w", err)
		}

		log.Println("[DEBUG] Bucket name:", *output.Name)

		for _, obj := range output.Contents {
			log.Println("[DEBUG] Object name:", *obj.Key)

			file, err := os.Create(*obj.Key)
			if err != nil {
				return fmt.Errorf("failed to create local file: %w", err)
			}

			downloader := s3manager.NewDownloaderWithClient(client.S3)
			_, err = downloader.Download(file, &s3.GetObjectInput{
				Bucket: aws.String(cfg.BucketName),
				Key:    obj.Key,
			})
			if err != nil {
				return fmt.Errorf("failed to download object file: %w", err)
			}

			log.Println("[DEBUG] Download successfully:", *obj.Key)
		}

		complete = output.IsTruncated != nil && !*output.IsTruncated
		token = output.NextContinuationToken
	}

	return nil
}

// Delete deletes any file from remote bucket.
func (client *Client) Delete(cfg *config.Config) error {
	var ok bool
	var err error

	if ok, err = objectHealthCheck(cfg.BucketName, cfg.ObjectKey, client); err != nil {
		return fmt.Errorf("object file health check failed: %w", err)
	}

	if !ok {
		log.Println("[DEBUG] Object file is not found.")
		return nil
	}

	_, err = client.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(cfg.ObjectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object file: %w", err)
	}

	log.Println("[DEBUG] Deletion successfully:", cfg.ObjectKey)

	return nil
}

// UploadWithPreSignedRequest requests a Presigned URL and uploads an object to a remote bucket.
func (client *Client) UploadWithPreSignedRequest(cfg *config.Config) error {
	req, _ := client.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(cfg.ObjectKey),
	})

	url, err := req.Presign(cfg.PreSignedUrlExpireLimit)
	if err != nil {
		return fmt.Errorf("failed to get presigning URL: %w", err)
	}

	log.Println("[DEBUG] Upload Presigned URL:", url)

	status, err := uploadWithPresignedUrl(url, cfg.UploadImagePath)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	log.Println("[DEBUG] Upload Presigned URL status:", status)

	return nil
}

// objectHealthCheck verifies the health check of an object file.
func objectHealthCheck(bucketName, objectKey string, client *Client) (bool, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	_, err := client.S3.HeadObject(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NotFound" {
			return false, nil
		}

		return false, fmt.Errorf("failed to get object metadata: %w", err)
	}

	return true, nil
}
