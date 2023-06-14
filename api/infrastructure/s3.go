package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	s3 *s3.S3
}

func NewS3Repository() (*S3Client, error) {
	s3Client := new(S3Client)
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	accessKey := "admin"
	secretKey := "password"
	region := "ap-northeast-1"
	endPoint := "http://127.0.0.1:9000"

	cfg := aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:           aws.String(region),
		Endpoint:         aws.String(endPoint),
		S3ForcePathStyle: aws.Bool(true),
	}

	s3Client.s3 = s3.New(s, &cfg)

	return s3Client, nil
}
