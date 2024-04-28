package config

import "time"

const (
	useMinIO                = true // MinIO を使用する場合は true に設定
	bucketName              = "sample-storage"
	objectKey               = "user01/sample.jpg"
	uploadFilePath          = "./test/sample.txt"
	uploadImagePath         = "./images/sample.jpg"
	accessKey               = "admin"
	secretKey               = "password"
	region                  = "ap-northeast-1"
	preSignedUrlExpireLimit = 1 * time.Minute
)

type Config struct {
	UseMinIO                bool
	BucketName              string
	ObjectKey               string
	UploadFilePath          string
	UploadImagePath         string
	AccessKey               string
	SecretKey               string
	Region                  string
	PreSignedUrlExpireLimit time.Duration
}

func Get() *Config {
	return &Config{
		UseMinIO:                useMinIO,
		BucketName:              bucketName,
		ObjectKey:               objectKey,
		UploadFilePath:          uploadFilePath,
		UploadImagePath:         uploadImagePath,
		AccessKey:               accessKey,
		SecretKey:               secretKey,
		Region:                  region,
		PreSignedUrlExpireLimit: preSignedUrlExpireLimit,
	}
}
