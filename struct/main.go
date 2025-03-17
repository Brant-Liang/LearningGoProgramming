package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type UserLogin struct {
	UserName        string
	Email           string
	Password        string
	ConfirmPassword string
	IsLoggedIn      bool
}

func (user *UserLogin) Login(userName string, password string, ConfirmPassword string) error {
	if user.UserName != userName {
		return errors.New("local user is not equal param of username")
	}
	if user.Password != password {
		return errors.New("local password is not equal param of password")
	}
	if user.ConfirmPassword != ConfirmPassword {
		return errors.New("local password is not equal param of confirm password")
	}
	user.IsLoggedIn = true
	return nil
}

func NewUserLogin() *UserLogin {
	return &UserLogin{
		UserName:        "admin",
		Email:           "admin@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
		IsLoggedIn:      false,
	}
}

func (user *UserLogin) Logout(userName string, password string, ConfirmPassword string) error {
	if user.UserName != userName {
		return errors.New("local user is not equal param of username")
	}
	if user.Password != password {
		return errors.New("local password is not equal param of password")
	}
	if user.ConfirmPassword != ConfirmPassword {
		return errors.New("local password is not equal param of confirm password")
	}
	user.IsLoggedIn = false
	return nil
}

func (user *UserLogin) updatePassword(password string) error {
	if user.Password == password {
		return errors.New("same password, please change it")
	}
	user.Password = password
	return nil
}

func main01() {
	user := NewUserLogin()
	err := user.Login("admin", "123456", "123456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("login success")
	}

	err1 := user.Logout("admin", "123456", "123456")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(user.IsLoggedIn)
		fmt.Println("logout success")
	}
	err2 := user.updatePassword("123457")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("update password success")
	}
}

type SizeOfStruct struct {
	a uint8
	b uint32
	c uint8
}

type SizeOfStructStr1 struct {
	a uint8  // 1
	b string // 16
	c uint8  // 1
}

type SizeOfStructStr2 struct {
	a bool   // 1
	b string // 16 字节 但最大对齐值为 8
	c string // 16 字节，8 字节对齐
}

func main02() {
	var a uint32 // 4
	var b uint64 // 8
	var c bool   // 1
	var d string // 16
	var e float32
	var f float64
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f))
}

func main03() {
	var s SizeOfStruct
	fmt.Printf("结构体 s 的地址: %p\n", &s)
	fmt.Printf("字段 a 的地址: %p\n", &s.a)
	fmt.Printf("字段 b 的地址: %p\n", &s.b)
	fmt.Printf("字段 c 的地址: %p\n", &s.c)
	fmt.Printf("结构体大小: %d 字节\n", unsafe.Sizeof(s))
}

func main04() {
	var s SizeOfStructStr1
	fmt.Printf("结构体 s 的地址: %p\n", &s)
	fmt.Printf("字段 a 的地址: %p\n", &s.a)
	fmt.Printf("字段 b 的地址: %p\n", &s.b)
	fmt.Printf("字段 c 的地址: %p\n", &s.c)
	fmt.Printf("结构体大小: %d 字节\n", unsafe.Sizeof(s))
}

type User struct {
	name string
	age  int
	sex  string
}

type UserInfo struct {
	user     User
	photo    string
	position string
}

func main05() {
	type empty struct{}
	a := empty{}                  // 作为通道传输信号
	fmt.Println(unsafe.Sizeof(a)) //0
}

type userInfo3 struct {
	Name string
	Age  int
	_    struct{}
}

func main() {
	a := userInfo3{Name: "aaa", Age: 18}
	fmt.Println(a)
	// 阻止初始化
	//b := userInfo3{"bbb", 18}。// 编译失败 too few values in MystKeyStruct{...}
	//fmt.Println(b)
}

type Set struct {
	items map[interface{}]interface{}
}
