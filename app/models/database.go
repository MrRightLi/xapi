package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var ORM *xorm.Engine

func init() {
	var err error
	ORM, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/serviceplat?charset=utf8")
	if err != nil {
		panic(err)
	}
}
