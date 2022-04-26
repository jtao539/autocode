package atom

import (
	"strings"
	"unicode"
)

// Case2Camel 下划线写法转为大驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func Case2SmallCamel(name string) string {
	return LowFirst(Case2SmallCamel(name))
}

// LowFirst 首字母小写
func LowFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
