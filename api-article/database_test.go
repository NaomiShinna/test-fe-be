package api_article

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// ...

// "database/sql"
// "time"

func OpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/article")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// test cek database

}

// db, err := sql.Open("mysql", "user:password@/dbname")
// if err != nil {
// 	panic(err)
// }
// // See "Important settings" section.
// db.SetConnMaxLifetime(time.Minute * 3)
// db.SetMaxOpenConns(10)
// db.SetMaxIdleConns(10)
