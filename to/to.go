// Package 序列化和反序列
package to

type To interface {
	// the v should be struct or map
	Encode(v interface{}) string
	// the v should be pointer or map
	Decode(data []byte, v interface{})
}

func Marshal(v interface{}, to To) string {
	return to.Encode(v)
}

func Unmarshal(data []byte, v interface{}, to To) {
	to.Decode(data, v)
}
