package main

import "strings"

var currentKey = "#current"
var rootKey = "#root"

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
