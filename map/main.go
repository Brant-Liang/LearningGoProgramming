package main

import (
	"fmt"
	"sync"
)

func main() {
	// make 关键字声明方式
	map1 := make(map[string]int)
	map1["a"] = 97
	map1["b"] = 98

	fmt.Println(map1["a"])
	fmt.Println(map1["b"])

	// 字面量
	map2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(map2["a"])
	fmt.Println(map2["b"])

	// map key 允许的类型
	// bool int string array 指针 结构体 接口
	// 不允许
	// slice maps 函数 channel
	// ❌ 错误: slice 作为 map key
	// sliceMap := map[[]int]string{} // 编译错误：invalid map key type []int

	// ❌ 错误: map 作为 map key
	// nestedMap := map[map[string]int]string{} // 编译错误：invalid map key type map[string]int

	// map 并发不安全 如果多个 goroutine 同时读写 map，会导致 数据竞争，可能引发 fatal error: concurrent map read and map write。
	m := make(map[int]int)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i * 10 // ❌ 这里会导致并发写入 map
		}(i)
	}

	wg.Wait()
	fmt.Println("Map:", m)
}
