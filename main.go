package main

import (
	"github.com/GotoRen/storage-controller/client"
	"github.com/GotoRen/storage-controller/config"
)

func main() {
	cfg := config.Get()

	// s3クライアントを作成
	client, err := client.NewS3Client(cfg)
	if err != nil {
		panic(err)
	}

	// 1. ローカルファイルをアップロード
	if err := client.Upload(cfg); err != nil {
		panic(err)
	}

	// // 2. オブジェクトファイルを読込
	// if err := client.Read(cfg); err != nil {
	// 	panic(err)
	// }

	// // 3. オブジェクトファイルをダウンロード
	// if err := client.Download(cfg); err != nil {
	// 	panic(err)
	// }

	// // 4. オブジェクトファイルを削除
	// if err := client.Delete(cfg); err != nil {
	// 	panic(err)
	// }
}
