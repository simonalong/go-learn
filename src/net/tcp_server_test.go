package net

import (
	"bufio"
	"fmt"
	"github.com/isyscore/isc-gobase/logger"
	"io"
	"net"
	"testing"
)

func TestTcpServer(t *testing.T) {
	// 监听本地的 8080 端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Info(err.Error())
	}
	defer listener.Close()

	logger.Info("Server is listening on port 8080...")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		// 处理连接，使用 goroutine 来处理每个连接
		go handleConnection(conn)
	}
}

func TestSendData(t *testing.T) {
	// 连接到服务端，假设服务端运行在本地的8080端口
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer conn.Close()

	// 创建一个bufio的Writer，方便写入数据
	writer := bufio.NewWriter(conn)

	// 要发送的数据
	data := "Hello, server!"

	// 发送数据
	_, err = writer.WriteString(data + "\n")
	if err != nil {
		logger.Error("Error writing to connection: %v", err.Error())
		return
	}
	writer.Flush()

	// 这里可以继续读取服务端的响应，如果需要的话
	// 例如：
	// response, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	log.Println("Error reading from connection:", err)
	// 	return
	// }
	// fmt.Println("Received from server:", response)

	fmt.Println("Data sent to server successfully.")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取数据
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				logger.Error(err.Error())
			}
			break
		}

		// 打印接收到的数据
		fmt.Printf("Received from client: %s\n", buffer[:n])

		// 回写数据给客户端
		_, err = conn.Write(buffer[:n])
		if err != nil {
			logger.Error(err.Error())
			break
		}
	}
}
