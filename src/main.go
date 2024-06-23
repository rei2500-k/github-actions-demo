package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	// GetのときTodoリストを返す
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		// DB接続情報
		const (
			host     = "demo_db"
			port     = 5432
			user     = "postgres"
			password = "postgres"
			dbname   = "postgres"
		)

		// DB接続
		conn := fmt.Sprintf(`
			host=%s
			port=%d
			user=%s
			password=%s
			dbname=%s
			sslmode=disable
		`, host, port, user, password, dbname)

		db, err := sql.Open("postgres", conn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// ToDoリストテーブルからレコードを取得
		rows, err := db.Query("select id, title, completed from todo_list")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var todos []*ToDo

		for rows.Next() {
			var id int
			var title string
			var completed bool

			err := rows.Scan(&id, &title, &completed)
			if err != nil {
				log.Fatalf("Error scanning row: %q", err)
			}

			// レコードを画面表示用に配列に入れる
			todos = append(todos, &ToDo{id, title, completed})
		}

		err = rows.Err()
		if err != nil {
			log.Fatalf("Error iterating over rows: %q", err)
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
