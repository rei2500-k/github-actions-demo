package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// DB接続
func ConnectDB(config DBConfig) (*sql.DB, error) {
	conn := fmt.Sprintf(`
		host=%s
		port=%d
		user=%s
		password=%s
		dbname=%s
		sslmode=disable
	`, config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("error open db connection: %w", err)
	}
	fmt.Println("connection success")
	return db, nil
}
