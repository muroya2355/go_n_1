package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

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

// User : 利用者データ
type Luser struct {
	ID    int
	Uname string
	Age   int
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
	// 借用者ID を文字列として格納する配列
	var userIDs []string
	// 利用者ID -> 利用者情報へのマップ
	userMap := make(map[int]Luser)

	// 書籍一覧の取得、格納
	rows, _ := Db.Query("SELECT * FROM book")
	for rows.Next() {
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.UserID)
		books = append(books, &book)

		// 借用者IDを文字列に変換して格納
		userIDs = append(userIDs, strconv.Itoa(book.UserID))
	}

	// SQL 文の組み立て
	sql := "SELECT id, uname, age FROM luser WHERE id IN (" + strings.Join(userIDs, ",") + ");"

	// 利用者テーブルから抽出、データの取得
	rows, _ = Db.Query(sql)
	for rows.Next() {
		var luser Luser
		err = rows.Scan(&luser.ID, &luser.Uname, &luser.Age)

		// 利用者データを map に格納
		userMap[luser.ID] = luser
	}

	for _, book := range books {
		// map から利用者データの取得
		book.Uname = userMap[book.UserID].Uname
		book.Age = userMap[book.UserID].Age
		// 結果の表示
		fmt.Println(book.ID, book.Title, book.Uname, book.Age)
	}

}
