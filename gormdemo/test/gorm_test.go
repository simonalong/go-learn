package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type NeoTable1 struct {
	Group string
	Name  string
	//UserName string
	//Age uint
	//Sort int
}

func TestConnect(t *testing.T) {
	//db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)

	//配置MySQL连接参数
	//username := "root"  //账号
	//password := "123456" //密码
	//host := "localhost" //数据库地址，可以是Ip或者域名
	//port := 23306 //数据库端口
	//Dbname := "envoy" //数据库名
	//timeout := "10s" //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/envoy?charset=utf8&parseTime=True&loc=Local", username, password, host, port)
	dsn := "root:123456@tcp(localhost:23306)/envoy?charset=utf8&parseTime=True&loc=Local"

	// 链接
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 增加
	db.Table("envoy_f_table1").Create(&NeoTable1{Group: "group", Name: "name"})

	// 删除
	//db.Delete(&NeoTable1{}).Where()

	// 修改

	// 查询：一行一列
	// 查询：一行多列
	// 查询：多行一列
	// 查询：多行多列
	// 查询：分页
	// 查询：存在否
}
