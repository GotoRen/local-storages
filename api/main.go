package main

import "github.com/GotoRen/local-storages/api/infrastructure"

func main() {
	bucketName := "test"
	objectKey := "sample.txt"
	uploadFilePath := "./sample.txt"

	// s3クライアントを作成
	client, err := infrastructure.NewS3Repository()
	if err != nil {
		panic(err)
	}

	// 1. ローカルファイルをアップロード
	if err := client.Upload(bucketName, objectKey, uploadFilePath); err != nil {
		panic(err)
	}

	// // 2. オブジェクトファイルを読込
	// if err := client.Read(bucketName, objectKey); err != nil {
	// 	panic(err)
	// }

	// // 3. オブジェクトファイルをダウンロード
	// if err := client.Download(bucketName); err != nil {
	// 	panic(err)
	// }

	// // 4. オブジェクトファイルを削除
	// if err := client.Delete(bucketName, objectKey); err != nil {
	// 	panic(err)
	// }
}
