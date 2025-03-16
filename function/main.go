package main

import (
	"errors"
	"fmt"
)

// go 语言支持多返回值
func funtionName() (float64, float64, error) {
	return 1.0, 2.0, errors.New("Something bad happened")
}

func swap(x, y uint64) (a, b uint64) {
	b, a = x, y
	return
}

func sum(nums ...uint64) uint64 {
	var total uint64
	for _, v := range nums {
		total += v
	}
	return total
}

func applyOperator(a, b int, operator func(int, int) int) int {
	return operator(a, b)
}

func main01() {
	add := func(x, y uint64) uint64 {
		return x + y
	}
	a, b := swap(1, 2)
	fmt.Println(a, b)
	fmt.Println(add(1, 2))

	counter := func() func() uint64 {
		count := 0
		return func() uint64 {
			count++
			return uint64(count)
		}
	}

	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}

// go 的访问修饰，首字母大写 首字母小写
func deferCallPanic() {
	defer func() {
		defer func() { fmt.Println("defer panic 之前 11") }()
		if err := recover(); err != nil {
			fmt.Println("recover panic:", err)
		}
	}()
	//defer func() { fmt.Println("defer panic 之前 11") }()
	defer func() { fmt.Println("defer panic 之前 22") }()
	panic("异常内容")
	defer func() { fmt.Println("defer panic 之后 33") }()
}

func main() {
	deferCallPanic()
	fmt.Println("end")
}
