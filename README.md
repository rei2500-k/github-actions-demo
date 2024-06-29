# github-actions-demo
GitHub Actions, Goを学ぶ目的でToDoリストを作成する。

## HTTPサーバ起動手順
1. コンテナ起動  
`docker compose up -d`

2. Goコンテナに入る  
`docker compose exec go bash`

3. HTTPサーバ起動  
`cd src`  
`go run main.go`

4. 接続確認  
ブラウザから[localhost:8080/](http://localhost:8080/)にアクセス

___

## DB接続
1. コンテナ起動  
`docker compose up -d`

2. DBコンテナに入る  
`docker compose exec db bash`

3. postgrsqlログイン  
`psql -U postgres`


## テスト
1. コンテナ起動  
`docker compose up -d`

2. Goコンテナに入る  
`docker compose exec db bash`

3. テスト実行ディレクトリに移動  
`cd src`

4. テスト実行  
`go test ./... -v`
