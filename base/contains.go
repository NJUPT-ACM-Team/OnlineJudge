package base

import "reflect"

func ToSlice(arrv interface{}) []interface{} {
	v := reflect.ValueOf(arrv)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func ArrayContains(arrv interface{}, v interface{}) bool {
	arr := ToSlice(arrv)
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}
