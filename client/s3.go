package client

import (
	"github.com/GotoRen/storage-api-sample/api"
	"github.com/GotoRen/storage-api-sample/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewS3Client returns s3 client.
func NewS3Client(cfg *config.Config) (*api.Client, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	s3Cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		Region:      aws.String(cfg.Region),
	}

	if cfg.UseMinIO {
		s3Cfg.S3ForcePathStyle = aws.Bool(true)
		s3Cfg.Endpoint = aws.String("http://localhost:9000")
	}

	return &api.Client{
		S3: s3.New(s, &s3Cfg),
	}, nil
}
