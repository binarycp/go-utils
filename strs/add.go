// Package a string utils
package strs

import "strings"

// 字符串拼接
func Concat(stirs ...string) string {
	var b strings.Builder
	for _, v := range stirs {
		//grow会发生内存逃逸
		//b.Grow(len(v))
		b.WriteString(v)
	}
	return b.String()
}
