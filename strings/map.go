// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-02-05 12:14

package strings

import (
	"log"
	"reflect"
)

// typeCmp to compare the tow map type and key, value type, must be the same
func typeCmp(f1, f2 interface{}) bool {
	t1, t2 := reflect.TypeOf(f1), reflect.TypeOf(f2)

	return t1.Kind() != reflect.Map || t2.Kind() != reflect.Map ||
		t1.Key() != t2.Key() || t1.Elem() != t2.Elem()
}

// MapMerge to merge mapB to mapA
func MapMerge(mapA, mapB interface{}) {
	va, vb := reflect.ValueOf(mapA), reflect.ValueOf(mapB)

	if typeCmp(mapA, mapB) {
		log.Printf("can not merge different map type: %v, %v", va.Type(), vb.Type())
		return
	}

	switch vb.Kind() {
	case reflect.Map:
		if vb.MapKeys() != nil {
			for _, v := range vb.MapKeys() {
				va.SetMapIndex(v, vb.MapIndex(v))
			}
		}
	default:
		return
	}
}
