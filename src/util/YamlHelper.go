package util

import (
	"strings"
)

type ConfigAppValue struct {
	Key       string
	Value     string
	Desc      string
	ValueType string
}

type CommonFlagEnum int8
type ConfigValueTypeEnum int8
type AllowPushEnum int8
type ConfigChangeTypeEnum int8

const (
	PRIVATE CommonFlagEnum = 0
	PUBLIC  CommonFlagEnum = 1
)

func YmlToConfigKeyList(key, yamlContent string) ([]ConfigAppValue, error) {
	propertiesStr, err := YamlToProperties(yamlContent)
	if err != nil {
		return nil, err
	}

	var configAppValueList []ConfigAppValue
	appValueList := KvPropertiesToKeyValueList(propertiesStr)
	for _, appValue := range appValueList {
		var configAppValue ConfigAppValue
		configAppValue.Key = key + "." + appValue.Key
		configAppValue.Value = appValue.Value
		configAppValue.ValueType = appValue.ValueType

		configAppValueList = append(configAppValueList, configAppValue)
	}

	return configAppValueList, nil
}

func KvPropertiesToKeyValueList(kvProperty string) []ConfigAppValue {
	if kvProperty == "" {
		return nil
	}

	var configAppValueList []ConfigAppValue
	if !strings.Contains(kvProperty, "=") {
		var configAppValue ConfigAppValue
		configAppValue.Key = ""
		configAppValue.Value = kvProperty
		configAppValue.ValueType = "STRING"

		configAppValueList = append(configAppValueList, configAppValue)
		return configAppValueList
	}

	var itemList = GetPropertiesItemLineList(kvProperty)
	for _, item := range itemList {
		if "" == item || strings.HasPrefix(item, "#") {
			continue
		}

		index := strings.Index(item, "=")
		var key = item[:index]
		var value = item[index+1:]

		var configAppValue ConfigAppValue
		configAppValue.Key = key
		configAppValue.Value = value
		configAppValue.ValueType = "STRING"
		configAppValueList = append(configAppValueList, configAppValue)
	}
	return configAppValueList
}

func KvPropertiesToKeyValueList1(kvProperty string) []ConfigAppValue {
	if kvProperty == "" {
		return nil
	}

	var configAppValueList []ConfigAppValue
	if !strings.Contains(kvProperty, "=") {
		var configAppValue ConfigAppValue
		configAppValue.Key = ""
		configAppValue.Value = kvProperty
		configAppValue.ValueType = "STRING"

		configAppValueList = append(configAppValueList, configAppValue)
		return configAppValueList
	}

	keys := ""
	values := ""

	var itemList = GetPropertiesItemLineList(kvProperty)
	for _, kvs := range itemList {
		if "" == kvs || strings.HasPrefix(kvs, "#") {
			continue
		}

		index := strings.Index(kvs, "=")
		if index < 0 {
			values += kvs
		} else {
			if strings.HasPrefix(kvs, " ") {
				values += kvs
				continue
			}

			if "" != keys {
				var configAppValue ConfigAppValue
				configAppValue.Key = keys
				configAppValue.Value = values
				configAppValue.ValueType = "STRING"
				configAppValueList = append(configAppValueList, configAppValue)

				keys = ""
				values = ""
			}
			keys += kvs[:index]
			values += kvs[index+1:]
		}
	}

	var configAppValue ConfigAppValue
	configAppValue.Key = keys
	configAppValue.Value = values
	configAppValue.ValueType = "STRING"
	configAppValueList = append(configAppValueList, configAppValue)
	return configAppValueList
}
