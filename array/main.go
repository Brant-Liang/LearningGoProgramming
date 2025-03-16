package main

import (
	"fmt"
	"unsafe"
)

// 数组： 一组数据的集合
// 数组长度：定长数组的长度是固定的
// 变长数组：不固定
// 数组中每个元素类型是相同的

// 切片
// 是一个动态数组，可以动态增减长度
// 切片是引用地址，它是对底层数据的引用，修改切片会影响底层数据
//

func arrayTest() {
	var arr [10]int
	arrStr := [3]string{"a", "b", "c"}

	fmt.Println(len(arr))

	fmt.Println(len(arrStr))

	for i := 0; i < len(arrStr); i++ {
		fmt.Println(arrStr[i])
	}

	for index, value := range arrStr {
		fmt.Println(index, value)
	}

	var matrix [3][3]int
	matrix[0][1] = 1

	fmt.Println("array size", unsafe.Sizeof(arr))
}

// 切片的创建
func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s[1:4])
	fmt.Println(s)
	s1 := s[1:2]
	sss := make([]int, 5, 10)
	fmt.Println("len", len(s1), "cap", cap(s1))
	fmt.Println("len", len(sss), "cap", cap(sss))

	fmt.Println("------")
	a := []int{1, 2, 3, 4, 5, 6, 7}
	a[0] = 100

	a = append(a, 10)
	fmt.Println(a)
}
