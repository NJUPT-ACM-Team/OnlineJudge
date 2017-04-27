package base

import (
	"strings"
)

func PBTagToFieldName(tag string) string {
	ret := ""
	s := strings.Split(tag, "_")
	for _, v := range s {
		ret += strings.ToUpper(v[0:1]) + v[1:]
	}
	return ret
}
