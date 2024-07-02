package tdengine

import (
	"database/sql"
	"fmt"
	"github.com/taosdata/driver-go/v3/af"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosSql"
	"testing"
)

func TestConnectOriginal(t *testing.T) {
	var taosDSN = "root:taosdata@tcp(localhost:6030)/"
	taos, err := sql.Open("taosSql", taosDSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err.Error())
		return
	}
	fmt.Println("Connected")
	defer taos.Close()
}

func TestConnectRest(t *testing.T) {
	var taosDSN = "root:taosdata@http(localhost:6041)/"
	taos, err := sql.Open("taosRestful", taosDSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err.Error())
		return
	}
	fmt.Println("Connected")
	defer taos.Close()
}

func TestConnect(t *testing.T) {
	conn, err := af.Open("localhost", "root", "taosdata", "", 6030)
	defer conn.Close()
	if err != nil {
		fmt.Println("failed to connect, err:", err)
	} else {
		fmt.Println("connected")
	}
}
