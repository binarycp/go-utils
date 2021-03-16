package slices

// 头部添加元素
// 加上第三个参数是为了避免内存逃逸
func UnShiftInt(ints []int, e int, ret []int) {
	//you should defined ret as follow:
	//ret := make([]int, len(ints)+1, len(ints)+1)
	copy(ret, []int{e})
	copy(ret[1:], ints)
}
