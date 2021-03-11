package slices

import (
	"fmt"
	"testing"
)

func TestDelAtInt(t *testing.T) {
	ints := make([]int, 0)
	ints = append(ints, []int{1, 2, 3, 4, 5}...)
	DelAtInt(ints, 4)
	ints = ints[:len(ints)-1]

	// output:
	// [1,2,3,4]
}

func TestDelAtMultiInt(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(DelAtMultiInt(ints, 3, 2))
}
