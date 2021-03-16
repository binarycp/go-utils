package slices

import (
	"fmt"
	"testing"
)

func TestUnShiftInt(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6}
	ret := make([]int, len(ints)+1, len(ints)+1)
	UnShiftInt(ints, 0, ret)
	fmt.Println(ret, ints)
}
