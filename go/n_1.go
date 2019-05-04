package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	ID     int
	Title  string
	UserID int
}
type Luser struct {
	ID    int
	Uname string
	Age   int
}

func main() {

	// Db: データベースに接続するためのハンドラ
	var Db *sql.DB
	// Dbの初期化
	Db, err := sql.Open("postgres", "host=postgres user=user password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// 書籍一覧の取得
	rows, _ := Db.Query("SELECT * FROM book")

	// 取得した各列について
	for rows.Next() {

		// 書籍データの取得
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.UserID)

		// 借用者IDをもとに、利用者テーブルから抽出、利用者データの取得
		var luser Luser
		Db.QueryRow("SELECT * FROM luser WHERE id = $1;", book.UserID).Scan(&luser.ID, &luser.Uname, &luser.Age)

		// 結果の表示
		fmt.Println(book.ID, book.Title, luser.Uname, luser.Age)
	}

}
