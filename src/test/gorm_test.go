package test

//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"testing"
//)
//
//func TestGorm(t *testing.T) {
//	//dsn := "neo_test:neo@Test123@tcp(127.0.0.1:3306)/demo1?charset=utf8mb4"
//	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//if nil != err {
//	//	return
//	//}
//	//fmt.Println(db)
//
//	dsn := "neo_test:neo@Test123@tcp(127.0.0.1:3306)/demo1?charset=utf8mb4"
//	db, err := gorm.Open(mysql.New(mysql.Config{
//		DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
//		DefaultStringSize: 256, // string 类型字段的默认长度
//		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
//		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
//		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
//		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
//
//
//		//DriverName: "mysql"
//		ServerVersion             string
//		DSN                       string
//		Conn                      gorm.ConnPool
//		SkipInitializeWithVersion bool
//		DefaultStringSize         uint
//		DefaultDatetimePrecision  *int
//		DisableDatetimePrecision  bool
//		DontSupportRenameIndex    bool
//		DontSupportRenameColumn   bool
//		DontSupportForShareClause bool
//	}))
//	if nil != err {
//		return
//	}
//	fmt.Println(db)
//}
