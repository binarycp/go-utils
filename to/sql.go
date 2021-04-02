package to

import (
	"reflect"
)

const (
	fieldTag = "\t"
	newTag   = "\n"
)

/*
resolve the string as follow
column1	column2
1	2
3	4
*/
type Sql struct{}

func (s Sql) Encode(v interface{}) string {
	of := reflect.TypeOf(v)
	switch of.Kind() {
	case reflect.Map:
		println("map")
	case reflect.Struct:

	}
	return ""
}

func (s Sql) Decode(data []byte, v interface{}) {
	panic("implement me")
}
