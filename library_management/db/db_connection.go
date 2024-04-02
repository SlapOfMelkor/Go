package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connectdb() {
	var err error
	postgres := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	Db, err = sql.Open("postgres", postgres)
	if err != nil {
		fmt.Println("Bağlantı sağlanamadı:", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Veritabanına ping atılamadı:", err)
		return
	}

	fmt.Println("Veritabanı bağlantısı başarıyla kuruldu.")
}
