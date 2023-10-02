# storage-api-sample

## 使い方

### 1. ローカル環境構築

```shell
### composeを起動
$ docker compose up -d
```

### 2. MinIO アクセス

- http://localhost:9001/login
- ユーザ：admin
- パスワード：password

### 3. バケットを作成

- バケット名：`test`

![picture](https://github.com/GotoRen/storage-api-sample/assets/63791288/7b560292-ff5a-453f-a5e1-82edfddddfa8)

### 4. コメントアウトを必要に応じてチェックアウト

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
```
