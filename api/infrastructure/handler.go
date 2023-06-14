package infrastructure

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *S3Client) Upload(bucket, object, uploadFile string) error {
	file, err := os.Open(uploadFile)
	if err != nil {
		return fmt.Errorf("failed to get local file: %w:", err)
	}
	defer file.Close()

	uploader := s3manager.NewUploaderWithClient(s.s3)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload local file: %w:", err)
	}

	fmt.Println("[DEBUG] Upload successful:", uploadFile)

	return nil
}

func (s *S3Client) Read(bucket, object string) error {
	objOutput, err := s.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return fmt.Errorf("failed to get object information: %w", err)
	}
	defer objOutput.Body.Close()

	bodyBytes, err := ioutil.ReadAll(objOutput.Body)
	if err != nil {
		return fmt.Errorf("failed to read object body: %w", err)
	}

	contents := string(bodyBytes)

	fmt.Print("[DEBUG] Object body: ", contents)

	return nil
}

func (s *S3Client) Download(bucket string) error {
	var token *string
	for complete := false; !complete; {
		in := s3.ListObjectsV2Input{Bucket: &bucket, ContinuationToken: token}

		out, err := s.s3.ListObjectsV2(&in)
		if err != nil {
			return fmt.Errorf("failed to get bucket information: %w", err)
		}

		fmt.Println("[DEBUG] Bucket name:", *out.Name)

		for _, obj := range out.Contents {
			fmt.Println("[DEBUG] Object name:", *obj.Key)

			file, err := os.Create(*obj.Key)
			if err != nil {
				return fmt.Errorf("failed to create local file: %w", err)
			}

			downloader := s3manager.NewDownloaderWithClient(s.s3)
			_, err = downloader.Download(file, &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    obj.Key,
			})
			if err != nil {
				return fmt.Errorf("failed to download object file: %w", err)
			}

			fmt.Println("[DEBUG] Download successful:", *obj.Key)
		}

		complete = out.IsTruncated != nil && !*out.IsTruncated
		token = out.NextContinuationToken
	}

	return nil
}

func (s *S3Client) Delete(bucket, object string) error {
	var ok bool
	var err error

	if ok, err = objectHealthCheck(bucket, object, s); err != nil {
		return fmt.Errorf("object file health check failed: %w", err)
	}

	if !ok {
		fmt.Println("[DEBUG] Object file is not found.")
		return nil
	}

	_, err = s.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object file: %w", err)
	}

	fmt.Println("[DEBUG] Deletion successful:", object)

	return nil
}

func objectHealthCheck(bucket, object string, s *S3Client) (bool, error) {
	in := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	}

	_, err := s.s3.HeadObject(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NotFound" {
			return false, nil
		}

		return false, fmt.Errorf("failed to get object metadata: %w", err)
	}

	return true, nil
}
