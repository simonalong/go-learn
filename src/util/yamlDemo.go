package main

import (
	"fmt"
	"github.com/magiconair/properties"
	yaml2 "github.com/simonalong/tools/yaml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	//var path_pre = "/Users/zhouzhenyong/project/private/Doramon/Ocean/src/test/resources/";
	//propertiesTest()
	//showStruct2(path_pre + "yml/base.yml")
	//showStruct(path_pre + "yml/base1.yml")
	//showStruct(path_pre + "yml/array1.yml")
	//showStruct(path_pre + "yml/array2.yml")
	//showStruct(path_pre + "yml/array3.yml")
	//showStruct(path_pre + "yml/array4.yml")
	//showStruct(path_pre + "yml/array5.yml")
	//showStruct(path_pre + "yml/array6.yml")
	//showStruct(path_pre + "yml/array7.yml")
	//showStruct(path_pre + "yml/cron.yml")
	//showStruct(path_pre + "yml/multi_line.yml")
	//showStruct(path_pre + "property/test.yml")
	//
	//p := properties.NewProperties()
	//p.SetProperty("nihao", "ok")
	//
	//propertiesTest()
	//showStruct2();
	jsonToYaml()
}

func jsonToYaml() {
	//str1 := "[]"
	str1 := "{\"appName\":\"reds\",\"configItemKey\":null,\"profile\":\"default\"}"
	fmt.Println(util.JsonToYaml(str1))
}

func showStruct2(path string) {
	//fmt.Println(path)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err.Error())
	}

	resultMap := make(map[string]interface{})
	fmt.Println(string(yamlFile))

	err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	//fmt.Println(resultMap)

	bytes2, err := yaml.Marshal(resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println(string(bytes2))
}

func showStruct(path string) {
	fmt.Println(path)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err.Error())
	}

	resultMap := make(map[string]interface{})
	fmt.Println(string(yamlFile))

	err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(resultMap)
}

func propertiesTest() {
	fmt.Println("======================================================================")
	var propertiesExample = []byte(`
p_id[0]=0001
p_id[1]=0002
p_type=donut
# 好的
p_name=Cake
p_ppu=0.55
p_batters.batter.type=Regular
`)
	//config := viper.New()
	//config.SetConfigType("properties")

	//r := bytes.NewReader(propertiesExample)
	//config.ReadConfig(r)

	pro := properties.NewProperties()
	pro.Load(propertiesExample, properties.UTF8)

	//fmt.Println(pro.String())
	//fmt.Println(pro.Map())
	fmt.Println(pro.Get("p_type"))
	fmt.Println(pro.Get("p_batters"))

	//err := pro.Load(r)
	//if nil != err {
	//	fmt.Println("异常")
	//}
	//
	////	var configjson = map[string]interface{}{}
	////	if err :=config.Unmarshal(&configjson);err !=nil{
	////		fmt.Println(err)
	////	}
	////
	////fmt.Println(configjson)
	//////r := bytes.NewW(propertiesExample)
	////
	//////
	//
	//value, exist := pro.Property("p_id")
	//if exist {
	//	fmt.Println("p_id ==== " + value)
	//} else {
	//	fmt.Println("p_id 不存在 ")
	//}
	//
	//
	//
	//
	//str := ""
	////bytes.NewBufferString(str)
	//
	//bf := bytes.NewBufferString(str)
	//
	//pro.Store(bf)
	//
	//fmt.Println("=============--------------=======")
	//fmt.Println(str)
	//fmt.Println(bf.String())

	//	config := viper.New()
	//	//config.AddConfigPath("./kafka_demo")
	//	config.Set("app[0]", "11")
	//	config.Set("app[1]", "12")
	//	config.Set("app[2]", "13")
	//	config.SetConfigName("config")
	//	config.SetConfigType("properties")
	//	//if err := config.ReadInConfig(); err != nil {
	//	//	panic(err)
	//	//}
	//	//fmt.Println(config.GetString("appId"))
	//	//fmt.Println(config.GetString("secret"))
	//	//fmt.Println(config.GetString("host.address"))
	//	//fmt.Println(config.GetString("host.port"))
	//
	//	//直接反序列化为Struct
	//	var configjson = map[string]interface{}{}
	//	if err :=config.Unmarshal(&configjson);err !=nil{
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(configjson)
	//	//fmt.Println(configjson["secret"])
	//	//fmt.Println(configjson["host"])
	//
	//	//config.WriteConfig()

	//pro.AllKeys()
	//config.WriteConfigAs()

	//
	//if v.properties == nil {
	//	v.properties = properties.NewProperties()
	//}
	//p := v.properties
	//for _, key := range v.AllKeys() {
	//	_, _, err := p.Set(key, v.GetString(key))
	//	if err != nil {
	//		return ConfigMarshalError{err}
	//	}
	//}
	//_, err := p.WriteComment(f, "#", properties.UTF8)
	//if err != nil {
	//	return ConfigMarshalError{err}
	//}
}

//func NewReader(b []byte) *io.Reader {
//	return &io.Reader{b, 0, -1}
//}
//
//func NewWriter(b []byte) *io.Writer {
//	return &io.Writer{b, 0, -1}
//}
