package main

import (
	"gopkg.in/yaml.v2"
	"log"
)

/**
 *  1.yaml <---> properties
 *  2.yaml <---> json
 *  3.yaml <---> map
 *  4.yaml <---> list
 *  5.yaml <---> kvList
 */

type KeyValue struct {
	Key   string
	Value interface{}
}

func YamlToPropertiesStr(contentOfYaml string) string {
	// yaml 到 map
	dataMap := YamlToMap(contentOfYaml)

}

//
//
//func PropertiesStrToYaml(contentOfProperties string) string {
//
//}
//

func YamlToMap(contentOfYaml string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(contentOfYaml), &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return resultMap
}

func MapToYaml(dataMap map[string]interface{}) string {
	bytes2, err := yaml.Marshal(dataMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return ""
	}
	return string(bytes2)
}

//
//func YamlToJson(contentOfYaml string) string {
//
//}
//
//func JsonToYaml(contentOfJson string) string {
//
//}
//
//func YamlToList(contentOfYaml string) []interface{} {
//
//}
//
//func KvListToYaml(kvList []KeyValue) string {
//
//}

// 进行深层嵌套的map数据处理
func MapToProperties(dataMap map[string]interface{}) {
	for key, value := range dataMap {
		// map的判断

		// 集合判断

		// 分片判断

		// string判断

		// 其他类型
	}
}

func doMapToProperties(value map[string]interface{}, prefix string) []string {

}
