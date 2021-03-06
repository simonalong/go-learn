package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type NeoMap map[string]interface{}

func getDb() (*sql.DB, error) {
	url := "neo_test:neo@Test123@tcp(localhost:3306)/neo?charset=utf8"
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("加载失败")
	}

	// 核查数据库
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("数据库链接失败")
	}

	fmt.Println("数据库链接成功")

	return db, nil
}

func main() {
	//sqlInsert();
	//sqlDelete();
	//sqlUpdate();
	sqlOne()
	//sqlList();
	//sqlValue();
	//sqlValues();
	//sqlPage();
	//sqlCount();
	//sqlExist();

	execute1()
	exe2()
}

func sqlInsert() {
	db, _ := getDb()
	rows, err := db.Query(`INSERT INTO neo_table1 (name) VALUES ("xys")`)
	//rows, err := db.Query(`select * from neo_table1`)
	defer rows.Close()
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}

	var result int
	rows.Scan(&result)
	log.Printf("insert result %v\n", result)
}

type neoTable1_tem struct {
	Id    int64
	Name  string
	Group string
}

func (table *neoTable1_tem) String() string {
	return fmt.Sprintf("NeoTable1: Id = %d , Name = %s, Group = %s\n", table.Id, table.Name, table.Group)
}

func sqlOne() {

}

func execute(rows *sql.Rows) {
	columns, _ := rows.Columns()
	columnsLength := len(columns)
	nums := make([]interface{}, columnsLength)
	point := make([]interface{}, columnsLength)
	for i := range nums {
		point[i] = &nums[i]
	}

	//result := map[string]interface{}
	for rows.Next() {
		err := rows.Scan(point...)
		if err != nil {
			log.Fatal(err)
		}
		for i := range nums {
			fmt.Println("数据：%s, %s", columns[i], nums[i])
			//result[columns[i]] = nums[i]
		}
	}
	//fmt.Println(result)
}

func execute1() (NeoMap, error) {
	db, _ := getDb()
	//rows, _:= db.Query("select `name`, `group` from neo_table1")
	stmt, _ := db.Prepare("select `name`, `group`, `id` from neo_table1 where id = ?")
	rows, _ := stmt.Query("1")
	columns, _ := rows.Columns()
	columnsLength := len(columns)
	nums := make([]interface{}, columnsLength)
	point := make([]interface{}, columnsLength)
	for i := range nums {
		point[i] = &nums[i]
	}
	result := make(map[string]interface{})
	for rows.Next() {
		var id int
		var name string
		var group string
		err := rows.Scan(point...)
		//err := rows.Scan(&name, &group, &id)
		if err != nil {
			log.Fatal(err)
			return nil, nil
		}
		//for i := range nums {
		//	//fmt.Println("=====", reflect.TypeOf(point[i]))
		//	//fmt.Println("=====", reflect.ValueOf(point[i]))
		//	//fmt.Println("=====", reflect.ValueOf(nums[i]))
		//	//fmt.Println("=====", reflect.Indirect(reflect.ValueOf(point[i])))
		//	//fmt.Println("=====", reflect.Indirect(reflect.ValueOf(nums[i])))
		//	//fmt.Println("=====", reflect.ValueOf(nums[i]))
		//
		//	result[columns[i]] = fmt.Sprintf("%s", nums[i])
		//}

		fmt.Println("数据：")
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(group)
	}

	fmt.Println(result)
	return result, nil
}

func exe2() {
	db, _ := getDb()
	stmt, _ := db.Prepare("select `name`, `group` from neo_table1 where id = ?")
	rows, _ := stmt.Query("1")
	execute(rows)
}
