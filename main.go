package main

import "fmt"

// := 不能用全局变量

// const 定义枚举 iota 只能和 const() 使用
const (
	BEIJING  = 10 * iota
	SHanghai // 10
	Shenzhen // 20
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 1, iota * 2
	i, k
)

func main() {
	var a int
	var b string
	var c = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("type of a = %T\n", a)
	fmt.Println("Hello World")

	const length int = 10
}
