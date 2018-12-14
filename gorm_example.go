package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type demoTable struct {
	ID int `gorm:"int;not null;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(20)"`
	Email string `gorm:"type:varchar(200)"`
}

func main() {
	println("--> connecting to mysql...")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/crm?autocommit=true&charset=utf8")
	if err != nil {
		panic(err)
	}

	println("--> opening debug")
	db.LogMode(true)
	
	println("--> creating table")
	if !db.HasTable(&demoTable{}) {
		db.CreateTable(&demoTable{})
	}

	println("--> inserting or updating data")
	tmp := demoTable{}
	db.Where(&demoTable{
		Name: "tang",
	}).Assign(&demoTable{
		Name: "tang",
		Email: "a@ie9.org",
	}).FirstOrCreate(&tmp)

	println("--> printing result")
	fmt.Printf("result: %v\n", tmp)

	println("--> closing database")
	db.Close()
}
