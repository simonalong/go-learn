package test

import (
	"encoding/base64"
	"fmt"
	"github.com/go-eden/routine"
	"testing"
	"time"
)

func TestBase64Demo1(t *testing.T) {
	// a2V5YQ==
	fmt.Println(Base64Encode("key1"))
	// dmFsdWVh
	fmt.Println(Base64Encode("valuea"))

	fmt.Println("===")
	// a2V5Yg==
	fmt.Println(Base64Encode("keyb"))
	// dmFsdWVh
	fmt.Println(Base64Encode("valuea"))

	fmt.Println("===")
	// a2V5Yw==
	fmt.Println(Base64Encode("key1"))
	// dmFsdWVh
	fmt.Println(Base64Encode("valuea"))
	// dmFsdWVi
	fmt.Println(Base64Encode("valueb"))

	fmt.Println("===")
}

func TestBase64Demo2(t *testing.T) {
	fmt.Println(Base64Decode("YmFy"))
}

func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64EncodeByte(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(dst string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(dst)
	if err != nil {
		return "", err
	}
	return string(src), nil
}

func TestBase64Demo4(t *testing.T) {
	go func() {
		time.Sleep(time.Second)
	}()
	goid := routine.Goid()
	goids := routine.AllGoids()
	fmt.Printf("curr goid: %d\n", goid)
	fmt.Printf("all goids: %v\n", goids)
}

var nameVar = routine.NewLocalStorage()

func TestBase64Demo5(t *testing.T) {
	nameVar.Set("hello world")
	fmt.Println("name: ", nameVar.Get())

	// other goroutine cannot read nameVar
	go func() {
		fmt.Println("name1: ", nameVar.Get())
	}()

	// but, the new goroutine could inherit/copy all local data from the current goroutine like this:
	routine.Go(func() {
		fmt.Println("name2: ", nameVar.Get())
	})

	// or, you could copy all local data manually
	ic := routine.BackupContext()
	go func() {
		routine.InheritContext(ic)
		fmt.Println("name3: ", nameVar.Get())
	}()

	time.Sleep(time.Second)
}
