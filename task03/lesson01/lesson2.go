package lesson01

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
	dsn = "root:root@tcp(127.0.0.1:3306)/qa?charset=utf8mb4&parseTime=True&loc=Local"
)

type employees struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

type book struct {
	Id     int
	Title  string
	Author string
	Price  float32
}

func Run() {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	/*var em employees
	err = db.Get(&em, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		fmt.Println("查询失败:", err)
	}
	fmt.Println(em)
	var em01 employees
	err = db.Get(&em01, "SELECT * FROM employees where department='技术部'  LIMIT 1")
	if err != nil {
		fmt.Println("查询失败:", err)
	}
	fmt.Println(em01)*/
	var bs []book
	err = db.Select(&bs, "select * from book where price>50")
	if err != nil {
		fmt.Println("查询失败:", err)
	}

	fmt.Println(bs)

}
