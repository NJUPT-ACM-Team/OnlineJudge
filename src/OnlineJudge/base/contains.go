package base

import "reflect"

func IsNilOrZero(x interface{}) bool {
	if x == nil {
		return true
	}
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func toSlice(arrv interface{}) []interface{} {
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
	if arrv == nil {
		return false
	}
	arr := toSlice(arrv)
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}
