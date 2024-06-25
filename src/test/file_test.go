package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/simonalong/gole/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"testing"
	goTime "time"
)

const BufferSize = 100

func TestReadFile(t *testing.T) {

	url := "http://localhost:31120/api/cloud/cockpit/buried/uploadFile2"
	path := "/Users/zhouzhenyong/project/go-private/go-learn/src/mqtt/emqx.go"
	req, err := NewFileUploadRequest(url, path)
	if err != nil {
		fmt.Printf("error to new upload file request:%s\n", err.Error())
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error to request to the server:%s\n", err.Error())
		return
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(body)
}

// NewFileUploadRequest ...
func NewFileUploadRequest(url, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	dataMap := map[string]interface{}{}
	dataMap["ok"] = "sdf"

	bytes, _ := json.Marshal(dataMap)
	payload := strings.NewReader(string(bytes))
	//

	// 文件写入 body
	writer := multipart.NewWriter(body)
	//writer.WriteField("ok", string(bytes))

	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "application/json")
	writer.CreatePart(h)
	body.ReadFrom(payload)

	part, err := writer.CreateFormFile("tem", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	// 其他参数列表写入 body
	for k, v := range dataMap {
		if err := writer.WriteField(k, util.ToString(v)); err != nil {
			continue
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}

type TestEntity struct {
	Age int
}

var osInfoFilePre = "./proc/run_info_"
var osInfoFilePathPre = "/Users/zhouzhenyong/tem/proc"

func TestGobaseFile(t *testing.T) {

	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = "sdfasd"

	//file.AppendFile(osInfoFilePre + time.TimeToStringFormat(time.Now(), time2.FmtYMd) + ".log", isc.ObjectToJson(dataMap))

	//dir_list, e := ioutil.ReadDir(osInfoFilePathPre)
	//if e != nil {
	//	fmt.Println("read dir error")
	//	return
	//}
	//for i, v := range dir_list {
	//	fmt.Println(i, "=", v.Name())
	//}

	//fmt.Println(fileOverdue("run_info_2022-03-10.log"))

	deleteThresholdFile()
}

func deleteThresholdFile() {
	fileList, err := ioutil.ReadDir("/Users/zhouzhenyong/tem/proc")
	if err != nil {
		fmt.Println("异常：文件路径：/Users/zhouzhenyong/tem/proc")
		return
	}
	for _, fileData := range fileList {
		if fileData.IsDir() {
			continue
		}

		// 文件过期删除
		if fileOverdue(fileData.Name()) {
			fmt.Println("文件过期删除：" + fileData.Name())
			os.Remove("/Users/zhouzhenyong/tem/proc/" + fileData.Name())
		}
	}
}

// 判断文件是否过期
func fileOverdue(fileName string) bool {
	lastIndex := strings.LastIndex(fileName, ".log")
	if lastIndex < 0 {
		return false
	}
	fileName = fileName[len("run_info_"):lastIndex]
	fileTime, err := goTime.ParseInLocation("2006-01-02", fileName, goTime.Local)
	if nil != err {
		fmt.Println("read dir error")
		return false
	}

	threshold := goTime.Now().AddDate(0, 0, -7)
	if fileTime.Before(threshold) {
		return true
	}

	return false
}

func TestCreateFile(t *testing.T) {
	exist, _ := PathExists("./tt/ok")
	if !exist {
		os.MkdirAll("./tt/ok", os.ModePerm)
	}
	os.Create("./tt/ok/demo.txt") // 如果文件已存在，会将文件清空。
	//fmt.Println(fp, err)  // &{0xc000076780} <nil>
	//fmt.Printf("%T", fp)  // *os.File 文件指针类型
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func TestFile(t *testing.T) {
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))
	//fmt.Println(rand.Int63n(100))

}

func getTableName(fileName string) string {
	if fileName == "" {
		return ""
	}

	lastIndex := strings.LastIndex(fileName, "_")
	if lastIndex <= -1 {
		return ""
	}

	return fileName[:lastIndex]
}
