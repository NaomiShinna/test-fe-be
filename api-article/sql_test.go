package api_article

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO post_article(title, content, category, created_date, updated_date, status) VALUES('Judul', 'isi konten', 'kategori', '2022-10-10 05:15:15', '2022-10-12 05:15:15', 'Publish')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert to post_article")

}

func TextReturnSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, title, content, category, created_date, updated_date, status FROM post_article"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, title, content, category, created_date, updated_date, status string
		err = rows.Scan(&id, &title, &content, &category, &created_date, &updated_date, &status)

		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("title: ", title)
		fmt.Println("content: ", content)
		fmt.Println("category: ", category)
		fmt.Println("created_date: ", created_date)
		fmt.Println("updated_date: ", updated_date)
		fmt.Println("status: ", status)
	}
}
