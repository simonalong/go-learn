//package main
//
//import (
//	"fmt"
//	mqtt "github.com/eclipse/paho.mqtt.golang"
//)

package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

func main2() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://10.30.30.78:11883").SetClientID("isc-config-service")

	opts.SetUsername("admin")
	opts.SetPassword("1Sysc0re!")
	opts.SetPingTimeout(1 * time.Second)
	opts.SetKeepAlive(60 * time.Second)

	// 消息的处理函数
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		var messageData = string(msg.Payload())
		fmt.Println("收到消息")
		fmt.Println(messageData)
	})

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("connect fail")
		return
	}

	// 订阅主题
	if token := c.Subscribe("nup/tenant/status", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println("create topic fail")
		return
	}

	//var broker = "10.30.30.78"
	//var port = 11883
	//opts := mqtt.NewClientOptions()
	//opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	//opts.SetClientID("isc-config-service")
	//opts.SetUsername("admin")
	//opts.SetPassword("1Sysc0re!")
	//opts.SetDefaultPublishHandler(messagePubHandler)
	//opts.OnConnect = connectHandler
	//opts.OnConnectionLost = connectLostHandler
	//client := mqtt.NewClient(opts)
	//if token := client.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}
	sub(c)
	publish(c)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main1() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://10.30.30.78:11883")
	opts.SetClientID("isc-config-service")
	opts.SetUsername("admin")
	opts.SetPassword("1Sysc0re!")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)
}

func sub(client mqtt.Client) {
	topic := "nup/tenant/status"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("nup/tenant/status", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func main() {
	// value
	//str := "ZQQazAg7bh7Jj36yd82GqQ=="
	// valuexxx
	str := "I2p2C0GnrOlJ/fEQEI65vw=="
	passwordKey := "iscConfigService"
	fmt.Println("-----start-----")
	fmt.Println(AesDecryptECB(str, passwordKey))
	fmt.Println("-----end-----")
}

func AesDecryptECB(content string, key string) string {
	b, _ := base64.StdEncoding.DecodeString(content)
	cipher, _ := aes.NewCipher([]byte(key))
	d := make([]byte, len(b))
	size := 16
	for bs, be := 0, size; bs < len(b); bs, be = bs+size, be+size {
		cipher.Decrypt(d[bs:be], b[bs:be])
	}
	return strings.TrimSpace(string(d))
}

func padding(src []byte) []byte {
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

func AesEncryptECB(content string, key string) string {
	b := padding([]byte(content))
	cipher, _ := aes.NewCipher([]byte(key))
	d := make([]byte, len(b))
	size := 16
	for bs, be := 0, size; bs < len(b); bs, be = bs+size, be+size {
		cipher.Encrypt(d[bs:be], b[bs:be])
	}
	return base64.StdEncoding.EncodeToString(d)
}
