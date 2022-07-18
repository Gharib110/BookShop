package mysql2

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var UserDB *sql.DB

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root", "dapperblondie1380", "127.0.0.1", "book_store")

	var err error
	UserDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
		return
	}

	if err = UserDB.Ping(); err != nil {
		panic(err)
		return
	}

	UserDB.SetConnMaxLifetime(time.Minute * 3)
	UserDB.SetMaxOpenConns(15)
	UserDB.SetMaxIdleConns(10)
	log.Println("We have connected to MYSQL")
}
