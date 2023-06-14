# local-storages

## 使い方

### 1. 環境構築

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

![picture](https://github.com/GotoRen/local-storages/assets/63791288/50ac2cd1-ec98-40f4-8d08-b7cf74b71772)

### 4. `api/main.go`の 1~4 のコメントアウトを必要に応じてチェックアウトしながら実行

```go
// 1. ローカルファイルをアップロード
if err := client.Upload(bucketName, objectName, uploadFilePath); err != nil {
	panic(err)
}

// 2. オブジェクトファイルを読込
if err := client.Read(bucketName, objectName); err != nil {
	panic(err)
}

// 3. オブジェクトファイルをダウンロード
if err := client.Download(bucketName); err != nil {
	panic(err)
}

// 4. オブジェクトファイルを削除
if err := client.Delete(bucketName, objectName); err != nil {
	panic(err)
}
```
