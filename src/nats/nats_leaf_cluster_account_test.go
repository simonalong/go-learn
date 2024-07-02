package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

// 订阅：来自于叶节点1，leaf1的证书
func TestLeafClusterAccountSubFromLeaf1(t *testing.T) {
	// Connect to a server
	//nc, err := nats.Connect("nats://127.0.0.1:4122")
	nc, err := nats.Connect("nats://127.0.0.1:4122", nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/seatak/account_iot/user_app.creds"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("leaf.>", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于主集群节点1，account_iot 账户的 user_app 用户
func TestLeafClusterAccountSubFromCluster1_iot(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/seatak/account_iot/user_app.creds"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("admin.*", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于主集群节点1，account_admin 账户的 user_admin 用户
func TestLeafClusterAccountSubFromCluster1_admin(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/seatak/account_admin/user_admin.creds"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("admin.*", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 发布：叶子节点1
func TestLeafClusterAccountPubOnLeaf1(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4122", nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/seatak/account_leaf1/leaf1_admin.creds"))

	nc.Publish("leaf.ok", []byte("Hello World from leaf1"))
	nc.Close()
}

// 发布：集群节点1，account_admin 账户的 user_admin 用户
func TestLeafClusterAccountPubOnServer1_admin(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4222", nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/seatak/account_admin/user_admin.creds"))

	nc.Publish("admin.test", []byte("Hello World from server1"))
	nc.Close()
}

// 发布：集群节点1，account_admin 账户的 user_admin 用户
func TestLeafClusterAccountPubOnServer1_iot(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4222", nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/seatak/account_iot/user_app.creds"))

	nc.Publish("leaf.test", []byte("Hello World from server1"))
	nc.Close()
}
