package config

const (
	bucketName     = "test"
	objectKey      = "sample.txt"
	uploadFilePath = "./test/sample.txt"

	accessKey = "admin"
	secretKey = "password"
	region    = "ap-northeast-1"
	endPoint  = "http://127.0.0.1:9000"
)

type Config struct {
	BucketName     string
	ObjectKey      string
	UploadFilePath string
	AccessKey      string
	SecretKey      string
	Region         string
	EndPoint       string
}

func Get() *Config {
	return &Config{
		BucketName:     bucketName,
		ObjectKey:      objectKey,
		UploadFilePath: uploadFilePath,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		Region:         region,
		EndPoint:       endPoint,
	}
}
