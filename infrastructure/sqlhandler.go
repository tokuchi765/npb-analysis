package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// SQLHandler SQLのコネクションをハンドリングする
type SQLHandler struct {
	GormDB *gorm.DB
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

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.GormDB = db

	return sqlHandler
}
