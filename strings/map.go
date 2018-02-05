// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-02-05 12:14

package strings

import (
	"log"
	"reflect"
)

// MapMerge to merge mapB to mapA
func MapMerge(mapA, mapB interface{}) {
	a, b := reflect.ValueOf(mapA), reflect.ValueOf(mapB)
	if a.Kind() != reflect.Map || b.Kind() != reflect.Map {
		log.Printf("type not map %s:%s, %s:%s", mapA, a.Kind(), mapB, b.Kind())
		return
	}
	switch b.Kind() {
	case reflect.Map:
		if b.MapKeys() != nil {
			for _, v := range b.MapKeys() {
				a.SetMapIndex(v, b.MapIndex(v))
			}
		}
	default:
		return
	}
}
