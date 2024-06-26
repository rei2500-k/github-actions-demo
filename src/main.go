package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rei2500-k/github-actions-demo/db"
	"github.com/rei2500-k/github-actions-demo/handlers"
)

func main() {
	// DB接続情報
	dbConfig := db.DBConfig{
		Host:     "demo_db",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
	}
	// DB接続
	db, err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal("%w", err)
	}
	defer db.Close()

	todoHandler := &handlers.TodoHandler{DB: db}

	// ルーティング設定
	r := mux.NewRouter()
	r.HandleFunc("/todos", todoHandler.GetList).Methods("GET")

	// HTTPサーバー起動
	http.ListenAndServe(":8080", r)
}
