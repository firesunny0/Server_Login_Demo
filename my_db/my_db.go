package my_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var AuthDb *sqlx.DB

func InitDb() (err error) {
	AuthDb, err = sqlx.Connect("mysql", "root:123@(localhost:3306)/users")
	if err != nil {
		fmt.Printf("connect DataBase failed, err:%v\n", err)
		return
	}
	AuthDb.SetMaxOpenConns(20)
	AuthDb.SetMaxIdleConns(10)
	fmt.Printf("connect DataBase success\n")
	return
}
