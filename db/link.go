package db

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

const (
	MaxConns int = 10
	MixConns int = 1
)

func Init (){
	db, err = sql.Open("mysql","root:guihuachu@tcp(127.0.0.1:3306)/gushi?charset=utf8&parseTime=true")
	if err != nil{
		panic(err)
	}
	db.SetMaxIdleConns(MaxConns)
	db.SetMaxOpenConns(MixConns)
	err = db.Ping()
	if err != nil{
		panic(err)
	}
}