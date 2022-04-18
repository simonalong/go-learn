package test

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"os"
	"sync"
	"testing"
	"xorm.io/core"
)

var instance *xorm.Engine
var once sync.Once

type ORMOperation func(session *xorm.Session) error

func GetDb() *xorm.Engine {
	//if nil != instance {
	//	 return instance
	//}
	once.Do(func() {
		// db:
		//  address: 10.30.30.78
		//  port: 23306
		//  user: isyscore
		//  password: Isysc0re
		// url格式：[username]:[password]@tcp([ip]:[port])/[database]?charset=utf8
		url := fmt.Sprintf("%s:%s@tcp(%s:%s)/isc_config?charset=utf8", "isyscore", "Isysc0re", "10.30.30.78", "23306")
		db, err := xorm.NewEngine("mysql", url)
		if err != nil {
			println(err.Error())
			instance = nil
		}
		instance = db
	})

	os.Getenv("")

	// 显示sql
	instance.ShowSQL(true)

	// 配置日志
	instance.Logger().SetLevel(core.LOG_DEBUG)
	return instance
}

func Tx(f ORMOperation) (err error) {
	session := GetDb().NewSession()
	err = session.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover rollback:%s\r\n", p)
			err := session.Rollback()
			if err != nil {
				return
			}
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Printf("error rollback:%s\r\n", err)
			err := session.Rollback()
			if err != nil {
				return
			} // err is non-nil; don't change it
		} else {
			err = session.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = f(session)
	return err
}

func TestExec(t *testing.T) {
	res, err := GetDb().Exec("select * from config_center_tenant where tenant_id=? for update", "system")
	if err != nil {
		fmt.Println(res)
	}
}
