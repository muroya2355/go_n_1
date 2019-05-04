package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

// Book : 書籍データ
type Book struct {
	ID     int
	Title  string
	UserID int
	Uname  string
	Age    int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// メイン関数
func main() {

	// Db: データベースに接続するためのハンドラ
	var Db *sql.DB
	// Dbの初期化
	Db, err := sql.Open("postgres", "host=postgres user=user password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var books []*Book

	// 時間計測開始
	start := time.Now()

	// 書籍テーブルと利用者テーブルを結合、取得
	rows, _ := Db.Query("SELECT book.id, book.title, luser.uname, luser.age FROM book JOIN luser ON book.user_id = luser.id")

	// 取得した各列について
	for rows.Next() {

		// 書籍、利用者データの取得
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.Uname, &book.Age)
		books = append(books, &book)
	}

	// 時間計測終了
	end := time.Now()

	// 結果の表示
	for i := 0; i < min(9, len(books)); i++ {
		book := books[i]
		fmt.Println(book.ID, book.Title, book.Uname, book.Age)
	}

	fmt.Printf("処理時間: %f秒\n", (end.Sub(start)).Seconds())

}
