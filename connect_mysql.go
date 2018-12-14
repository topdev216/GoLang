package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func dbselect(db *sql.DB) {
	rows, err := db.Query("select dealer_code,dealer_name_short from dealer_base")
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	var dealer_code string
	var dealer_name_short string
	for rows.Next() {
		err := rows.Scan(&dealer_code, &dealer_name_short)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(dealer_code, dealer_name_short)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cbbweb?charset=utf8")
	if err != nil {
		log.Fatalln("open database failed")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	dbselect(db)
}
