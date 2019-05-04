package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func main() {

	// Db: データベースに接続するためのハンドラ
	var Db *sql.DB
	// Dbの初期化
	Db, err := sql.Open("postgres", "host=postgres user=user password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// データ件数
	const nBook = 200000
	const nLuser = 10000

	fmt.Println("書籍データの追加中...")
	for i := 5; i < nBook; i++ {

		rand.Seed(time.Now().UnixNano())

		// 書籍データをテーブルに追加
		Db.Exec("INSERT INTO book (title, user_id) VALUES ('" + RandString(3) + "', " + strconv.Itoa(rand.Intn(nLuser)+1) + ")")
	}

	fmt.Println("利用者データの追加中...")
	for i := 5; i < nLuser; i++ {

		rand.Seed(time.Now().UnixNano())

		// 利用者データをテーブルに追加
		Db.Exec("INSERT INTO luser (uname, age) VALUES ('" + RandString(2) + " " + RandString(3) + "', " + strconv.Itoa(rand.Intn(70)+10) + ")")
	}

}
