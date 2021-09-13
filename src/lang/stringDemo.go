package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//str := "21203129"
	//fmt.Println(str)
	//fmt.Println(strings.TrimSpace(str))
	//fmt.Println(strings.Trim(str, " "))
	//fmt.Println(strings.Replace(str, "2", "8", 2))
	//
	//fmt.Println(strings.Index(str, "="))
	//index := strings.Index(str, "o")
	//fmt.Println(str[:index])
	//fmt.Println(str[index:])
	//
	//ints := []int{23, 1, 23, 15, 2, 312, 42}
	//sort.Ints(ints)
	//
	//fmt.Println(ints)
	//
	//ints = []int{}
	//fmt.Println(ints)

	//for i, data := range str {
	//	r, _ := strconv.Atoi(string(data))
	//	fmt.Println(i, " ", r)
	//}

	//fmt.Println(str[:len(str)-1])

	//i := interp.New(interp.Options{})
	//
	//i.Use(stdlib.Symbols)
	//
	//i.Use(interp.Exports{
	//	"a": {
	//		"val": reflect.ValueOf(int64(11)),
	//	},
	//})
	//
	//
	//
	//_, err := i.Eval(`import "fmt"`)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = i.Eval(`fmt.Println(a.val)`)
	//if err != nil {
	//	panic(err)
	//}

	test1()

	// 97 ~ 122
	//data := "a.z"
	// 65~90
	//data := "AZ中"
	//
	//for _, i2 := range data {
	//	fmt.Println(i2)
	//}
}

var currentKey = "#current"
var rootKey = "#root"

func test1() {
	errMsg := "名字#root.Age不合法，还有当前的值#current不满足要求，还有#root.Name is not Empty"
	fmt.Println(errMsg)
	fmt.Println(errMsgChange(errMsg))
}

func errMsgChange(errMsg string) string {
	var matchKeys []string
	var chgMsg strings.Builder
	chgMsg.WriteString("sprintf(\"")

	var b strings.Builder
	b.Grow(len(errMsg))

	matchIndex := 0
	matchLength := 0
	for infoIndex, data := range errMsg {
		c := string(data)
		if c == "#" {
			if findCurrentKey(infoIndex, 0, errMsg) {
				matchIndex = 0
				matchLength = len(currentKey)
				b.WriteString("%v")
				matchKeys = append(matchKeys, "current")
				continue
			} else if find, size, wordKey := findRootKey(infoIndex, 0, errMsg); find {
				matchIndex = 0
				matchLength = size
				b.WriteString("%v")
				matchKeys = append(matchKeys, "root"+wordKey)
				continue
			}
		} else if matchIndex+1 < matchLength {
			matchIndex++
			continue
		} else {
			b.WriteString(c)
		}
	}

	chgMsg.WriteString(b.String())
	chgMsg.WriteString("\", ")

	matchKeysSize := len(matchKeys)
	for i, data := range matchKeys {
		if i+1 < matchKeysSize {
			chgMsg.WriteString(data)
			chgMsg.WriteString(", ")
		} else {
			chgMsg.WriteString(data)
		}
	}
	chgMsg.WriteString(")")

	return chgMsg.String()
}

func findCurrentKey(infoIndex, matchIndex int, info string) bool {
	if matchIndex >= len(currentKey) {
		return true
	}
	if info[infoIndex:infoIndex+1] == currentKey[matchIndex:matchIndex+1] {
		return findCurrentKey(infoIndex+1, matchIndex+1, info)
	}
	return false
}

func findRootKey(infoIndex, matchIndex int, info string) (bool, int, string) {
	if matchIndex >= len(rootKey) {
		nextKeyLength := nextMatchKeyLength(info[infoIndex:])
		if nextKeyLength > 0 {
			return true, len(rootKey) + nextKeyLength, info[infoIndex : infoIndex+nextKeyLength]
		}
		return false, 0, ""
	}
	if info[infoIndex:infoIndex+1] == rootKey[matchIndex:matchIndex+1] {
		return findRootKey(infoIndex+1, matchIndex+1, info)
	}
	return false, 0, ""
}

// 下一个英文的单词长度
// 97 ~ 122
// 65 ~ 90
func nextMatchKeyLength(errMsg string) int {
	spaceIndex := strings.Index(strings.TrimSpace(errMsg), " ")
	toMatchMsg := errMsg
	if spaceIndex > 0 {
		toMatchMsg = errMsg[:spaceIndex]
	}
	var index = 0
	for _, c := range toMatchMsg {
		// 判断是否是英文字符：a~z、A~Z和点号"."
		if (c >= 97 && c <= 122) || (c >= 65 && c <= 90) || c == 46 {
			index++
			continue
		} else {
			return index
		}
	}
	return index
}

// 将其中的root.xx和current生成对应的占位符和sprintf字段，比如：数据#root.Age的名字#current不合法，转换为：sprintf("数据%v的名字%v不合法", root.Age, current)
//func errMsgToTemplate(errMsg string) string {
//	strings.Index(errMsg, "#root")
//}

func Replace(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		b.WriteString(s[start:j])
		b.WriteString(new)
		start = j + len(old)
	}
	b.WriteString(s[start:])
	return b.String()
}
