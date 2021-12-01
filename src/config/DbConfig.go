package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

var instance *xorm.Engine
var once sync.Once

func GetDb() *xorm.Engine {
	once.Do(func() {
		// url格式：[username]:[password]@tcp([ip]:[port])/[database]?charset=utf8
		url := "neo_test:neo@Test123@tcp(localhost:3306)/neo?charset=utf8"
		db, err := xorm.NewEngine("mysql", url)
		if err != nil {
			println(err.Error())
			instance = nil
		}
		instance = db
	})
	instance.ShowSQL(true)
	return instance
}
