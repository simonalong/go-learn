package main

import (
	"fmt"
	"github.com/SimonAlong/go-learn/src/config"
	"github.com/go-xorm/xorm"
)

var db = config.GetDb()

type NeoTable1 struct {
	Id    int64
	Name  string
	Group string
}

type NeoTable2 struct {
	Id    int64
	Name  string
	Group string
}

type Count struct {
	count int64
}

type ORMOperation func(session *xorm.Session) error

func (table *NeoTable1) String() string {
	return fmt.Sprintf("NeoTable1: Id = %d , Name = %s, Group = %s\n", table.Id, table.Name, table.Group)
}

func main() {
	truncate()
	db.ShowSQL(true)

	//insert()
	//save()
	//delete()
	//update()
	one()
	//one2()
	//one3()
	//list()
	//page()
	//value()
	//values()
	//count()
	//exist()
	//join()
	//tx()
	//batchInsert()
	//batchUpdate()
	//execOne()
	//execList()
	//execValue()
	//execValues()
}

func truncate() {
	exec, err := db.Exec("truncate neo_table1")
	if err != nil {
		return
	}

	println(exec)
}

func insert() {
	println("============ 测试 insert ============")
	data := NeoTable1{}
	data.Name = "insert_name1"
	data.Group = "insert_group1"

	// 这样不会获取生成的id
	// result, err := db.InsertOne(data)
	// 这样可以获取生成的id
	// INSERT INTO `neo_table1` (`name`,`group`) VALUES (?, ?) []interface {}{"name1", "group1"}
	result, err := db.InsertOne(&data)
	if err != nil {
		println(err.Error())
	}

	println(result)
	// 可以拿到自动生成的id
	println(data.Id)
}

func save() {
	println("============ 测试 save ============")
}

func delete() {
	println("============ 测试 delete ============")
	data := NeoTable1{}

	i, err := db.Where("`name`=? and `group`=?", "insert_name1", "insert_group1").Delete(data)
	if err != nil {
		println(err.Error())
		return
	}
	println(i)
}

func update() {
	println("============ 测试 update ============")
	insert()
	data := NeoTable1{}
	data.Name = "update_name"
	data.Group = "update_group"
	db.Where("`name`=? and `group`=?", "insert_name1", "insert_group1").Update(data)
}

func one() {
	println("============ 测试 one ============")
	db.Insert(&NeoTable1{Name: "update_name", Group: "update_group"})

	data := NeoTable1{}
	get, err := db.Where("`name`=? and `group`=?", "update_name", "update_group").Get(&data)
	if err != nil {
		println(err.Error())
		return
	}

	print(get)
	println(data.String())
}

func one2() {
	println("============ 测试 one2 ============")
	data := NeoTable1{}
	// 选择对应的列
	get, err := db.Select("Id").Where("`name`=? and `group`=?", "update_name", "update_group").Get(&data)
	if err != nil {
		println(err.Error())
		return
	}

	print(get)
	println(data.String())
}

// 单条数据用Get
func one3() {
	println("============ 测试 one3 ============")
	data := NeoTable1{}
	// 选择对应的列
	// 选择某个表的所有列：db.Select("table.*")
	// 选择某个表的几个列：db.Select("table.name, table.group")
	// 选择所有表的所有列：db.Select("*")
	get, err := db.Select("neo_table1.name, neo_table1.group").Where("`name`=? and `group`=?", "update_name", "update_group").Get(&data)
	if err != nil {
		println(err.Error())
		return
	}

	print(get)
	println(data.String())
}

// 多条数据用Find
func list() {
	println("============ 测试 list ============")

	datas := []NeoTable1{}

	db.Insert(&NeoTable1{Name: "in_name", Group: "in_group1"})
	db.Insert(&NeoTable1{Name: "in_name", Group: "in_group2"})

	err := db.Where("`name`=?", "in_name").Find(&datas)
	if err != nil {
		println(err.Error())
		return
	}

	for i := range datas {
		println(datas[i].String())
	}

	truncate()
}

// 分页这里主要的就是Limit函数
func page() {
	println("============ 测试 page ============")

	db.Insert(&NeoTable1{Name: "page_name", Group: "in_group1"})
	db.Insert(&NeoTable1{Name: "page_name", Group: "in_group2"})
	db.Insert(&NeoTable1{Name: "page_name", Group: "in_group3"})
	db.Insert(&NeoTable1{Name: "page_name", Group: "in_group4"})

	datas := []NeoTable1{}
	err := db.Where("name=?", "page_name").Limit(3, 0).Find(&datas)
	if err != nil {
		println(err.Error())
		return
	}

	for i := range datas {
		println(datas[i].String())
	}
}

func value() {
	println("============ 测试 value ============")

	db.Insert(&NeoTable1{Name: "value_name", Group: "value_group1"})
	db.Insert(&NeoTable1{Name: "value_name", Group: "value_group2"})

	data := NeoTable1{}
	_, err := db.Select("`group`").Where("name=?", "value_name").Get(&data)
	if err != nil {
		println(err.Error())
		return
	}

	println(data.String())
}

func values() {
	println("============ 测试 values ============")
	println("============ 测试 value ============")

	db.Insert(&NeoTable1{Name: "value_name", Group: "value_group1"})
	db.Insert(&NeoTable1{Name: "value_name", Group: "value_group2"})

	datas := []NeoTable1{}
	err := db.Select("`group`").Where("name=?", "value_name").Find(&datas)
	if err != nil {
		println(err.Error())
		return
	}

	for i := range datas {
		println(datas[i].String())
	}
}

func tx() {
	println("============ 测试 tx ============")
	session := db.NewSession()

	// 开启事务
	err := session.Begin()
	if err != nil {
		return
	}

	// 函数执行完毕后执行
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover rollback:%s\r\n", p)
			err := session.Rollback()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			panic(p)
		} else if err != nil {
			fmt.Printf("err2 rollback:%s\r\n", err)
			err := session.Rollback()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			err = session.Commit()
		}
	}()

	// 下面执行业务代码
	if _, err2 := session.Insert(&NeoTable1{Id: 200, Name: "value_name", Group: "value_group1"}); err2 != nil {
		err = err2
		return
	}

	if _, err3 := session.Insert(&NeoTable1{Id: 201, Name: "value_name", Group: "value_group2"}); err3 != nil {
		err = err3
		return
	}

	if _, err4 := session.Insert(&NeoTable1{Id: 202, Name: "value_name", Group: "value_group3"}); err4 != nil {
		err = err4
		return
	}
}

func count() {
	println("============ 测试 count ============")

	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group1"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})
	db.Insert(&NeoTable1{Name: "count_name", Group: "value_group2"})

	i, err := db.Table("neo_table1").Where("name=?", "count_name").Count()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 返回的个数
	println(i)
}

func exist() {
	println("============ 测试 exist ============")

	b, err := db.Table("neo_table1").Where("name=?", "count_name").Exist()
	if err != nil {
		println(err.Error())
		return
	}
	// true或者false
	println(b)
}

func join() {
	println("============ 测试 join ============")

	db.Insert(&NeoTable1{Name: "join_name", Group: "value_group1"})
	db.Insert(&NeoTable1{Name: "join_name", Group: "value_group2"})
	db.Insert(&NeoTable2{Name: "join_name", Group: "value_group1"})
	db.Insert(&NeoTable2{Name: "join_name", Group: "value_group2"})
	db.Insert(&NeoTable2{Name: "join_name", Group: "value_group3"})
	db.Insert(&NeoTable2{Name: "join_name", Group: "value_group4"})

	datas := []NeoTable1{}
	db.Table("neo_table1").Join("INNER", "neo_table2", "neo_table2.name = neo_table1.name").Find(&datas)

	for i := range datas {
		println(datas[i].String())
	}

	// 问题，就是怎么获取多个数据字段的类型呢，假设一个BO类型，这里面有多个表中的字段，这个怎么进行获取
}

func batchInsert() {
	println("============ 测试 batchInsert ============")

	datas := []NeoTable1{}[:]
	datas = append(datas, NeoTable1{Name: "batch_insert_name1", Group: "batch_insert_group1"})
	datas = append(datas, NeoTable1{Name: "batch_insert_name2", Group: "batch_insert_group2"})
	datas = append(datas, NeoTable1{Name: "batch_insert_name3", Group: "batch_insert_group3"})

	i, err := db.Insert(datas)
	if err != nil {
		println(err.Error())
		return
	}
	println(i)
}

func batchUpdate() {
	println("============ 测试 batchUpdate ============")
}

func execOne() {
	println("============ 测试 execOne ============")
}

func execList() {
	println("============ 测试 execList ============")
}

func execValue() {
	println("============ 测试 execValue ============")
}

func execValues() {
	println("============ 测试 execValues ============")

	data := make([]NeoTable1, 2)
	data[0].Name = "batch_name1"
	data[0].Group = "batch_group1"
	data[1].Name = "batch_name2"
	data[1].Group = "batch_group2"

	i, err := db.Insert(data)
	if err != nil {
		return
	}

	println(i)
}
