package main

import (
	"fmt"
	"github.com/lunny/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int64
	Name string
}

func main() {
	orm, err := xorm.NewEngine("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	defer orm.Close()

	orm.ShowSQL = true
	err = orm.CreateTables(&User{})
	if err != nil {
		panic(err)
	}

	_, err = orm.Insert(&User{1, "tang"})
	if err != nil {
		panic(err)
	}

	usr := User{}
	_, err = orm.Id(1).Get(&usr)
	if err != nil {
		panic(err)
	}
}
