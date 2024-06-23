# github-actions-demo
GitHub Actions, Goを学ぶ目的でToDoリストを作成する。

## HTTPサーバ起動手順
1. コンテナ起動  
`docker compose up -d`

2. コンテナに入る  
`docker compose exec go bash`

3. HTTPサーバ起動  
`cd src`  
`go run main.go`

4. 接続確認  
ブラウザから[localhost:8080/](http://localhost:8080/)にアクセス

___
