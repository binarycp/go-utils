package slices

import "testing"

func TestEqualToByte(t *testing.T) {
	a1 := []byte(`hello`)
	b1 := []byte(`hello`)

	var a2 []byte
	var b2 []byte

	a3 := []byte{0x16}
	b3 := []byte{0x18}

	println(EqualToByte(a1, b1))
	println(EqualToByte(a1, b2))
	println(EqualToByte(a2, b2))
	println(EqualToByte(a2, b1))
	println(EqualToByte(a3, b3))
}
