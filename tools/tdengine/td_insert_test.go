package tdengine

import (
	"database/sql"
	"fmt"
	"github.com/taosdata/driver-go/v3/af"
	"github.com/taosdata/driver-go/v3/common"
	"github.com/taosdata/driver-go/v3/common/param"
	_ "github.com/taosdata/driver-go/v3/taosSql"
	"testing"
	"time"
)

// sql插入
func TestTdInsert(t *testing.T) {
	var taosDSN = "root:taosdata@tcp(localhost:6030)/"
	taos, err := sql.Open("taosSql", taosDSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err.Error())
		return
	}

	sql := `INSERT INTO tdlearn.device1 USING tdlearn.meters TAGS('California.SanFrancisco', 2) VALUES ('2023-10-03 14:38:05.000', 10.30000, 219, 0.31000) ('2023-10-03 14:38:15.000', 12.60000, 218, 0.33000) ('2023-10-03 14:38:16.800', 12.30000, 221, 0.31000)
	tdlearn.device1 USING tdlearn.meters TAGS('California.SanFrancisco', 3) VALUES ('2023-10-03 14:38:16.650', 10.30000, 218, 0.25000)`
	result, err := taos.Exec(sql)
	if err != nil {
		fmt.Println("failed to insert, err:", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("failed to get affected rows, err:", err)
	}

	fmt.Println(rowsAffected)
}

// 参数绑定：插入；目前这个只能使用原生链接

func TestTdInsert2(t *testing.T) {
	conn, err := af.Open("localhost", "root", "taosdata", "", 6030)
	checkErr(err, "fail to connect")
	defer conn.Close()
	//prepareStable(conn)
	// create stmt
	stmt := conn.InsertStmt()
	defer stmt.Close()
	err = stmt.Prepare("INSERT INTO ? USING tdlearn.meters TAGS(?, ?) VALUES(?, ?, ?, ?)")
	checkErr(err, "failed to create prepare statement")

	// bind table name and tags
	tagParams := param.NewParam(2).AddBinary([]byte("California.SanFrancisco")).AddInt(2)
	err = stmt.SetTableNameWithTags("tdlearn.device1", tagParams)
	checkErr(err, "failed to execute SetTableNameWithTags")

	// specify ColumnType
	var bindType *param.ColumnType = param.NewColumnType(4).AddTimestamp().AddFloat().AddInt().AddFloat()

	// bind values. note: can only bind one row each time.
	valueParams := []*param.Param{
		//param.NewParam(1).AddTimestamp(time.Unix(1648432611, 249300000), common.PrecisionMilliSecond),
		param.NewParam(1).AddTimestamp(time.Now(), common.PrecisionMilliSecond),
		param.NewParam(1).AddFloat(1022.3),
		param.NewParam(1).AddInt(21229),
		param.NewParam(1).AddFloat(0.31111),
	}
	err = stmt.BindParam(valueParams, bindType)
	checkErr(err, "BindParam error")
	err = stmt.AddBatch()
	checkErr(err, "AddBatch error")

	// bind one more row
	valueParams = []*param.Param{
		//param.NewParam(1).AddTimestamp(time.Unix(1648432611, 749300000), common.PrecisionMilliSecond),
		param.NewParam(1).AddTimestamp(time.Now(), common.PrecisionMilliSecond),
		param.NewParam(1).AddFloat(1212.6),
		param.NewParam(1).AddInt(21832),
		param.NewParam(1).AddFloat(0.333334),
	}
	err = stmt.BindParam(valueParams, bindType)
	checkErr(err, "BindParam error")
	err = stmt.AddBatch()
	checkErr(err, "AddBatch error")
	//execute
	err = stmt.Execute()
	checkErr(err, "Execute batch error")
}

func checkErr(err error, prompt string) {
	if err != nil {
		fmt.Printf("%s\n", prompt)
		panic(err)
	}
}

func prepareStable(conn *af.Connector) {
	_, err := conn.Exec("CREATE DATABASE power")
	checkErr(err, "failed to create database")
	_, err = conn.Exec("CREATE STABLE power.meters (ts TIMESTAMP, current FLOAT, voltage INT, phase FLOAT) TAGS (location BINARY(64), groupId INT)")
	checkErr(err, "failed to create stable")
	_, err = conn.Exec("USE power")
	checkErr(err, "failed to change database")
}
