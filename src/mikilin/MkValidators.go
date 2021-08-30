package mikilin

import (
	"fmt"
	"reflect"
	"strings"
)

type FieldMatcher struct {
}

type MatcherCollector func(string2 string)

type MatcherUnit struct {
	name             string
	matcherCollector MatcherCollector
}

var matchers []MatcherUnit

var fieldMap = make(map[string]map[string][]FieldMatcher)

func Check(object interface{}) (bool, error) {
	objType := reflect.TypeOf(object)
	fmt.Println(objType.String())
	for index, num := 0, objType.NumField(); index < num; index++ {
		subCondition := strings.Split(objType.Field(index).Tag.Get(MATCHER), ";")
		for _, subStr := range subCondition {
			subStr = strings.TrimSpace(subStr)
			buildMatcher(subStr)
		}
	}
	return true, nil
}

func buildMatcher(subCondition string) {
	for _, matcher := range matchers {
		matcher.matcherCollector(subCondition)
	}
}

func init() {
	/* 搜集匹配后的操作参数 */
	matchers = append(matchers, MatcherUnit{ERR_MSG, collectErrMsg})
	matchers = append(matchers, MatcherUnit{CHANGE_TO, collectChangeTo})
	matchers = append(matchers, MatcherUnit{ACCEPT, collectAccept})
	matchers = append(matchers, MatcherUnit{DISABLE, collectDisable})

	/* 构造匹配器 */
	matchers = append(matchers, MatcherUnit{VALUE, buildValuesMatcher})
	matchers = append(matchers, MatcherUnit{IS_NIL, buildIsNilMatcher})
	matchers = append(matchers, MatcherUnit{IS_BLANK, buildIsBlankMatcher})
	matchers = append(matchers, MatcherUnit{RANGE, buildRangeMatcher})
	matchers = append(matchers, MatcherUnit{MODEL, buildModelMatcher})
	matchers = append(matchers, MatcherUnit{ENUM_TYPE, buildEnumTypeMatcher})
	matchers = append(matchers, MatcherUnit{CONDITION, buildConditionMatcher})
	matchers = append(matchers, MatcherUnit{CUSTOMIZE, buildCustomizeMatcher})
	matchers = append(matchers, MatcherUnit{REGEX, buildRegexMatcher})
}

func collectErrMsg(subCondition string) {

}

func collectChangeTo(subCondition string) {

}

func collectAccept(subCondition string) {

}

func collectDisable(subCondition string) {

}

func buildValuesMatcher(subCondition string) {
	if !strings.Contains(subCondition, VALUE) || !strings.Contains(subCondition, EQUAL) {
		return
	}

	index := strings.Index(subCondition, "=")
	value := subCondition[index+1:]
	data := []interface{}{value}
	// todo
	fmt.Println(data)
}

func buildIsNilMatcher(subCondition string) {

}

func buildIsBlankMatcher(subCondition string) {

}

func buildRangeMatcher(subCondition string) {

}

func buildModelMatcher(subCondition string) {

}

func buildEnumTypeMatcher(subCondition string) {

}

func buildConditionMatcher(subCondition string) {

}

func buildRegexMatcher(subCondition string) {

}

func buildCustomizeMatcher(subCondition string) {

}
