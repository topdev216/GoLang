package main

import (
	"fmt"
	"ie9.org/database"
	"log"
	"strings"
)

func main() {
	fmt.Println("请指定源数据库连接（车巴巴），例如：")
	fmt.Println("\t远程：root:123456@tcp(127.0.0.1:3306)/database?charset=utf8")
	fmt.Println("\t本机：root:123456@/database")
	fmt.Print("请输入：")

	var conn string
	_, err := fmt.Scanln(&conn)
	if err != nil {
		log.Fatalln("请指定数据库连接！")
	}

	if database.Connect(conn) != 0 {
		log.Fatalln("数据库连接失败！")
	} else {
		log.Println("数据库连接成功！")
	}

	fmt.Println("\n请选择要同步的表：")
	fmt.Println("\t1. car_series")
	fmt.Println("\t2. car_model")
	fmt.Println("\t3. activity")
	fmt.Println("\t4. promotions")
	fmt.Println("\t5. dealer_base")
	fmt.Println("\t6. financial")
	fmt.Println("\t7. financial_brand")
	fmt.Println("\t8. financial_series")
	fmt.Println("\t9. financial_dealer")
	fmt.Println("\t10. dealer_series")
	fmt.Println("\t11. dealer_price")
	fmt.Print("用逗号将序号隔开，不要带空格：")

	var t string
	_, err = fmt.Scanln(&t)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range strings.Split(t, ",") {
		if v == "1" {
			go database.Car_series()
		}
		if v == "2" {
			go database.Car_model()
		}
		if v == "3" {
			go database.Activity()
		}
		if v == "4" {
			go database.Promotions()
		}
		if v == "5" {
			go database.Dealer_base()
		}
		if v == "6" {
			go database.Financial()
		}
		if v == "7" {
			go database.Financial_brand()
		}
		if v == "8" {
			go database.Financial_series()
		}
		if v == "9" {
			go database.Financial_dealer()
		}
		if v == "10" {
			go database.Dealer_series()
		}
		if v == "11" {
			go database.Dealer_price()
		}
	}

}
