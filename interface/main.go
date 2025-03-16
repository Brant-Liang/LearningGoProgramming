package main

import "fmt"

func updateValue(n *int) {
	*n = 100
}

func returnValue() *int {
	x := 10
	return &x
}
func main01() {
	x := 10
	updateValue(&x)
	var p *int
	p = &x
	fmt.Println(*p)
	var pp *int
	if p != nil {
		fmt.Println(*pp)
	} else {
		fmt.Println("nil")
	}

	aa := returnValue()
	fmt.Println(aa)
}

type Animal interface {
	Species()
	Bark()
	Eat()
}

/*
修饰变量时，可以是任意类型
修饰方法时，是声明方法的集合
所有的对象都需要实现 interface 里声明的所有方法
*/
type Dog struct {
	DogType  string
	Language string
	Food     string
}

func (d *Dog) Bark() {
	fmt.Println(d.Language)
}

func NewDog() (*Dog, error) {
	return &Dog{
		DogType:  "golden retriever",
		Language: "English",
		Food:     "bone",
	}, nil
}

func (d *Dog) Species() {
	fmt.Println(d.DogType)
}

func (d *Dog) Eat() {
	fmt.Println("Eating", d.Food)
}

func main02() {
	var a interface{}
	var b interface{}
	// 空接口
	a = 1
	b = 2
	// 比较动态类型和动态值
	fmt.Println(a == b) // false
}

func main() {
	var aldog Animal
	aldog = new(Dog)

	aldog.Eat()
	aldog.Bark()
}
