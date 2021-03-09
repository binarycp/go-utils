package strs

import "testing"

func TestConcat(t *testing.T) {
	println(Concat("1", "2", "3"))
	println(Concat("我", "100", "分"))
	println(Concat("foo", "100", "分"))
}
