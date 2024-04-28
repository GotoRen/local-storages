package main

import (
	"log"

	"github.com/GotoRen/storage-api-sample/client"
	"github.com/GotoRen/storage-api-sample/config"
)

func main() {
	cfg := config.Get()

	// s3クライアントを作成
	client, err := client.NewS3Client(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// // 1. ローカルファイルをアップロード
	// if err := client.Upload(cfg); err != nil {
	// 	log.Fatal(err)
	// }

	// // 2. オブジェクトファイルを読込
	// if err := client.Read(cfg); err != nil {
	// 	log.Fatal(err)
	// }

	// // 3. オブジェクトファイルをダウンロード
	// if err := client.Download(cfg); err != nil {
	// 	log.Fatal(err)
	// }

	// // 4. オブジェクトファイルを削除
	// if err := client.Delete(cfg); err != nil {
	// 	log.Fatal(err)
	// }

	// 5. Presigned URL を使用してファイルをアップロード
	if err := client.UploadWithPreSignedRequest(cfg); err != nil {
		log.Fatal(err)
	}
}
