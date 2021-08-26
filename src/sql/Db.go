package main

import "github.com/SimonAlong/go-learn/src/config"

var db = config.GetDb()

type NeoTable1 struct {
	Id    int64
	Name  string
	Group string
}

func main() {
	insert()
	save()
	delete()
	update()
	one()
	list()
	page()
	value()
	values()
	count()
	exist()
	join()
	tx()
	batchInsert()
	batchUpdate()
	execOne()
	execList()
	execValue()
	execValues()
}

func insert() {
	println("============ 测试 insert ============")
	db.ShowSQL(true)
	data := NeoTable1{}
	data.Name = "name1"
	data.Group = "group1"

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
}

func update() {
	println("============ 测试 update ============")
}

func one() {
	println("============ 测试 one ============")
}

func list() {
	println("============ 测试 list ============")
}

func page() {
	println("============ 测试 page ============")
}

func value() {
	println("============ 测试 value ============")
}

func values() {
	println("============ 测试 values ============")
}

func tx() {
	println("============ 测试 tx ============")
}

func count() {
	println("============ 测试 count ============")
}

func exist() {
	println("============ 测试 exist ============")
}

func join() {
	println("============ 测试 join ============")
}

func batchInsert() {
	println("============ 测试 batchInsert ============")
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
}
