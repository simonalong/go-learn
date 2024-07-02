package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

// 订阅：来自于叶节点1
func TestLeafClusterSubFromLeaf1(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4122")
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于叶节点2
func TestLeafClusterSubFromLeaf2(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4123")
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于主集群节点1
func TestLeafClusterSubFromCluster1(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于主集群节点2
func TestLeafClusterSubFromCluster2(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4223")
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于主集群节点3
func TestLeafClusterSubFromCluster3(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4224")
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 订阅：来自于主集群 所有节点
func TestLeafClusterSubFromCluster_all(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4222,nats://127.0.0.1:4223,nats://127.0.0.1:4224")
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

// 发布：叶子节点1
func TestLeafClusterPubOnLeaf1(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4122")

	nc.Publish("foo", []byte("Hello World from leaf1"))
	nc.Close()
}

// 发布：叶子节点2
func TestLeafClusterPubOnLeaf2(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4123")

	nc.Publish("foo", []byte("Hello World from leaf2"))
	nc.Close()
}

// 发布：集群节点1
func TestLeafClusterPubOnServer1(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4222")

	nc.Publish("foo", []byte("Hello World from server1"))
	nc.Close()
}

// 发布：集群节点2
func TestLeafClusterPubOnServer2(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4223")

	nc.Publish("foo", []byte("Hello World from server2"))
	nc.Close()
}

// 发布：集群节点3
func TestLeafClusterPubOnServer3(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4224")

	nc.Publish("foo", []byte("Hello World from server3"))
	nc.Close()
}

// 发布：集群 所有节点
func TestLeafClusterPubOnServer_all(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4222,nats://127.0.0.1:4223,nats://127.0.0.1:4224")

	nc.Publish("foo", []byte("Hello World from server-all"))
	nc.Close()
}
