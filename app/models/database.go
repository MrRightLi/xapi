package models

import (
	"io/ioutil"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xapi/config"
	"gopkg.in/yaml.v2"
	"log"
)

var ORM *xorm.Engine

func init() {
	conf := new(config.Yaml)
	yamlFile, err := ioutil.ReadFile("config/test.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	startMysql(conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Name)
}

func startMysql(root string, password string, host string, port string, db string)  {
	//ORM, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/serviceplat?charset=utf8")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", root, password, host, port, db)
	log.Println(dataSourceName)
	var err error
	ORM, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
