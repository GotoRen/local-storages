# storage-controller

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

![picture](https://github.com/GotoRen/local-storages/assets/63791288/50ac2cd1-ec98-40f4-8d08-b7cf74b71772)

### 4. コメントアウトを必要に応じてチェックアウト

```go
// 1. ローカルファイルをアップロード
if err := client.Upload(cfg); err != nil {
	panic(err)
}

// 2. オブジェクトファイルを読込
if err := client.Read(cfg); err != nil {
	panic(err)
}

// 3. オブジェクトファイルをダウンロード
if err := client.Download(cfg); err != nil {
	panic(err)
}

// 4. オブジェクトファイルを削除
if err := client.Delete(cfg); err != nil {
	panic(err)
}
```
