package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func do_get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func okcoin_ticker(coin string) map[string]interface{} {
	var data map[string]interface{}
	var result map[string]interface{}
	resp := do_get("https://www.okcoin.cn/api/v1/ticker.do?symbol=" + coin)
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		panic(err)
	}

	result = data["ticker"].(map[string]interface{})
	result["date"] = data["date"].(interface{})
	return result
}

func init_sqlite(db_path string) *sql.DB {
	_, err := os.Stat(db_path)
	db_exists := err == nil || os.IsExist(err)

	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}

	if db_exists {
		return db
	}

	sqlStmt := `CREATE TABLE ticker (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		date INTEGER NOT NULL,
		buy REAL NOT NULL,
		sell REAL NOT NULL,
		high REAL NOT NULL,
		low REAL NOT NULL,
		last REAL NOT NULL,
		vol REAL NOT NULL
	)`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	return db
}

func insert_ticker(db *sql.DB, data map[string]interface{}) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare("INSERT INTO ticker (date, buy, sell, high, low, last, vol) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(data["date"], data["buy"], data["sell"], data["high"], data["low"], data["last"], data["vol"])
	tx.Commit()

	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	log.Println(string(result))
}

func main() {
	db := init_sqlite("okcoin.db")
	data := okcoin_ticker("btc_cny")
	insert_ticker(db, data)
	db.Close()
}
