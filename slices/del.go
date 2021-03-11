package slices

//移除指定位置的元素
//不会内存逃逸，但是需要调用方自己修改长度
//元素假如是指针类型，会导致内存泄漏
func DelAtInt(ints []int, index int) {
	ints = append(ints[:index], ints[index+1:]...)
}

//从指定位置移除指定数量的元素
//会发生内存逃逸，但是于此同时避免了内存泄漏
func DelAtMultiInt(ints []int, begin int, length int) []int {
	l := len(ints)
	ret := make([]int, begin, l-length)
	copy(ret, ints[:begin])
	return append(ret, ints[begin+length:]...)
}
