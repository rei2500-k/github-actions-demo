package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rei2500-k/github-actions-demo/db"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positives", args: args{a: 1, b: 2}, want: 3},
		{name: "negatives", args: args{a: -1, b: -2}, want: -3},
		{name: "zero", args: args{a: 1, b: 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoHandler_GetList(t *testing.T) {
	dbConfig := db.DBConfig{
		Host:     "demo_db",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
	}
	db, _ := db.ConnectDB(dbConfig)
	h := &TodoHandler{DB: db}
	req, _ := http.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(h.GetList)
	handler.ServeHTTP(rec, req)

	if got := rec.Code; got != http.StatusOK {
		t.Errorf("GetList() = %v, want %v", got, http.StatusOK)
	}

	var todos []ToDo
	json.NewDecoder(rec.Body).Decode(&todos)

	if len(todos) != 1 {
		t.Errorf("GetList().len() = %v, want %v", len(todos), 1)
	}

	if todos[0].Title != "test" {
		t.Errorf("todo_list Title = %v, want %v", todos[0].Title, "test")
	}
}
