package api_article

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/article")
	if err != nil {
		panic(err)
	}

	// test cek database
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 & time.Minute)

	return db
}
