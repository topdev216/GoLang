package models

import (
	"fmt"
	"github.com/lunny/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/goconfig"
	"os"
)

type Database struct {
	Driver   string
	Username string
	Password string
	Url      string
}

type Admin struct {
	Id       int64
	Username string
	Password string
	Level    int
}

var (
	CONFIG_FILE string = "website.cfg"
)

func sqliteEngine() (*xorm.Engine, error) {
	// load db config
	db, err := loadDatabase()
	if err != nil {
		return nil, err
	}

	// set engine pointer
	return xorm.NewEngine(db.Driver, db.Url)
}

func loadDatabase() (*Database, error) {
	// load config file
	config, err := goconfig.LoadConfigFile(CONFIG_FILE)
	if err != nil {
		return nil, err
	}

	// read values out
	// modify here to support other db: mysql, mongodb, oracle...
	var db *Database
	db.Driver, err = config.GetValue("Database", "Driver")
	db.Url, err = config.GetValue("Database", "Url")
	fmt.Println(db.Driver, db.Username, db.Password, db.Url)
	return db, err
}

var admin *Admin = &Admin{}

func initAdmin(engine *xorm.Engine) error {
	// delete existing db file
	// load db config
	db, err := loadDatabase()
	if err != nil {
		return err
	}
	os.Remove(db.Url)

	// create table
	err = engine.CreateTables(admin)
	if err != nil {
		return err
	}

	// defer close
	defer engine.Close()

	// init admin user
	sum, err := engine.Insert(&Admin{
		Username: "admin",
		Password: "admin",
		Level:    0})
	if sum > 0 && err != nil {
		return err
	}
	return nil
}

func InitTables() error {
	orm, err := sqliteEngine()
	if err != nil {
		return err
	}
	return initAdmin(orm)
}
