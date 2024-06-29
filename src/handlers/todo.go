package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/lib/pq"
)

type ToDo struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	Completed bool        `json:"completed"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt pq.NullTime `json:"deleted_at"`
}

type TodoHandler struct {
	DB *sql.DB
}

func (h *TodoHandler) GetList(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(`
		select
			id, title, completed,
			created_at, updated_at, deleted_at
		from
			todo_list
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []ToDo{}
	for rows.Next() {
		var todo ToDo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func Add(a int, b int) int {
	return a + b
}
