package main

import (
	"encoding/json"
	"net/http"
)

type ToDo struct {
	ID        string
	Title     string
	Completed bool
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	// GetのときTodoリストを返す
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		// ダミーTodoリスト
		todos := []ToDo{
			{ID: "1", Title: "Learn Go", Completed: false},
			{ID: "2", Title: "Buy desk chair", Completed: false},
		}

		json.NewEncoder(w).Encode(todos)

		return
	}
}

func main() {
	// ルーティング設定
	http.HandleFunc("/", listTodos)

	// HTTPサーバー起動
	http.ListenAndServe(":8080", nil)
}
