# storage-api-sample

## 使い方

### 1. ローカル環境構築

```shell
### composeを起動
$ docker compose up -d
```

### 2. MinIO アクセス

- http://localhost:9001/login
- `AWS_ACCESS_KEY_ID`：admin
- `AWS_SECRET_ACESS_key`：password
- `REGION_NAME`：ap-northeast-1

### 3. バケットを確認

- バケット名：`sample-storage`

![image](https://github.com/GotoRen/storage-api-sample/assets/63791288/d576122e-fa29-4386-81fb-e3788e47832c)

### 3. 必要に応じてコメントインして実行

```go
// 1. ローカルファイルをアップロード
if err := client.Upload(cfg); err != nil {
	log.Fatal(err)
}

// 2. オブジェクトファイルを読込
if err := client.Read(cfg); err != nil {
	log.Fatal(err)
}

// 3. オブジェクトファイルをダウンロード
if err := client.Download(cfg); err != nil {
	log.Fatal(err)
}

// 4. オブジェクトファイルを削除
if err := client.Delete(cfg); err != nil {
	log.Fatal(err)
}

// 5. Presigned URL を使用してファイルをアップロード
if err := client.UploadWithPreSignedRequest(cfg); err != nil {
	log.Fatal(err)
}
```
