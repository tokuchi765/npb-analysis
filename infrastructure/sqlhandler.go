package infrastructure

import (
	"database/sql"
	"fmt"
	"time"
)

// SQLHandler SQLのコネクションをハンドリングする
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler SQLHandlerを生成
func NewSQLHandler() *SQLHandler {
	conn, err := sql.Open("postgres", "host=localhost port=5555 password=postgres user=npb-analysis dbname=npb-analysis sslmode=disable")

	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		fmt.Println(err)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}
