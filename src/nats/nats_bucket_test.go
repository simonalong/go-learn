package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
)

func TestBucket(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	// 创建桶对象
	objectStore, _ := js.CreateObjectStore(&nats.ObjectStoreConfig{
		Bucket: "bucket_demo1",
	})

	// 删除桶对象
	//js.DeleteObjectStore("bucket_demo1")

	filePath1 := "/Users/zhouzhenyong/project/private-go/go-learn/main.go"
	filePath2 := "/Users/zhouzhenyong/project/private-go/go-learn/README.md"
	// 添加文件
	_, err := objectStore.PutFile(filePath1)
	_, err = objectStore.PutFile(filePath2)
	if err != nil {
		fmt.Println(err)
	}

	// 获取文件
	fileList, _ := objectStore.List()
	for _, fileInfo := range fileList {
		fmt.Println(fileInfo.Name)
	}

	objectStore.GetFile("/Users/zhouzhenyong/project/private-go/go-learn/main.go", "/Users/zhouzhenyong/project/private-go/go-learn/main1.go")
}
