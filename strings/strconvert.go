// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-01-30 16:14

// CamelToSnake can convert camelcase to snake, support string and []string types now

package strings

import (
	"reflect"
	"fmt"
	"strings"
	"unicode"
)

//CamelToSnake to convert string or []string to snake
func CamelToSnake(i interface{}) (interface{}, error) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		s := doConvert(i.(string))
		return s, nil
	case reflect.Slice:
		var out []string
		for idx := 0; idx < v.Len(); idx++ {
			s := doConvert(v.Index(idx).Interface().(string))
			out = append(out, s)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("err: %s type can not convert", v.Kind())
	}
}

//doConvert to convert camelcase to snake
func doConvert(s string) string {
	if s == strings.ToLower(s) {
		return s
	}
	var out []rune
	canConvert := func(i int) bool {
		if i > 0 {
			return 'A' <= s[i] && s[i] <= 'Z' && unicode.IsLower(rune(s[i-1])) && rune(s[i-1]) != '_'
		}
		return false
	}
	for i, v := range s {
		if canConvert(i) {
			out = append(out, '_')
		}
		out = append(out, v)
	}
	return strings.ToLower(string(out))
}
