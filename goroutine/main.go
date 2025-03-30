package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
进程、线程和协程
unix-c fork 进程： 0-3G的内存区域，所有进程是独享 0-3G的内存区域，独立程序的执行单元
0-3G: 用户态
3-4G: 内核空间

线程：进程中的一个执行单元，共享 0-3G 的资源

进程间的通信方式（IPC - Inner Process Communication）
- socket

线程： 进程中的一个执行单元，共享 0-3G 的资源
- socket
- 消息队列
- 全局变量

协程：用户态，由用户自己调度
GMP ---> 用户主动调度
- channel
- 全局变量
*/

// 子 gorutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Gorutine: i = %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

// 主 gorutine
func main01() {
	// 创建一个 gorutine
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main Gorutine: i = %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func speakEnglish(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("speak English")
}

func speakChinese(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("说中文")
}

func speakSpanish(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("speak Spanish")
}

func main02() {
	// go 运行时间是 M 的调度模型
	//go func() {
	//	fmt.Println("goroutine")
	//}()

	// 协程优雅退出： WaitGroup
	var wg sync.WaitGroup
	wg.Add(3)
	go speakEnglish(&wg)
	go speakChinese(&wg)
	go speakSpanish(&wg)
	wg.Wait() //等待所有 gorutine 结束
}

// make 和 new 区别
func main03() {

	// new 适合任意类型
	// 返回指针
	// 只分配 零值化的内存，但不初始化数据。
	ptr := new(int)
	fmt.Println(*ptr) // 0

	// make 仅适用于 slice、map、channel
	// 分配内存且初始化 返回值为初始化的值
	m := make(map[string]int)
	m["score"] = 100
	fmt.Println(m) // 输出：map[score:100]

}

func Add(a, b int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时，减少 WaitGroup 计数
	c <- a + b
}

func main04() {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(1)
	go Add(19, 8, c, &wg)
	fmt.Println(<-c)
	wg.Wait() // 等待所有 goroutine 结束
}

// goroutine 访问共享数据时，可能会导致数据竞争， Mutex 和 RWMutex

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main05() {
	var wg sync.WaitGroup

	counter := new(Counter)

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Value())
}

func main06() {
	var mu sync.Mutex         //互斥锁
	cond := sync.NewCond(&mu) //  goroutine 之间的同步和通信，返回一个 *sync.Cond 实例，可以用来等待和通知。
	ready := false
	go func() {
		time.Sleep(time.Second * 2)
		mu.Lock()
		ready = true
		mu.Unlock()
		cond.Signal()
	}()
	mu.Lock()
	for !ready {
		cond.Wait() // 会让 goroutine 暂停，直到 Signal() 或 Broadcast() 发生。

	}
	fmt.Println("condition met, proceeding")
	mu.Unlock()
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "start", j)
		results <- j * 2
		fmt.Println("worker", id, "end", j*2)
	}
}

func main07() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j < 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a < 5; a++ {
		fmt.Println(<-results)
	}
	close(results)
}

// context 进程之间传递上下文，这里是超时机制
func main() {
	crx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("received from channel") // ⛔️ 这里的"timeout"实际是通道收到数据时执行的
	case <-crx.Done():
		fmt.Println("crx done") // ⛔️ 这里是超时后执行的
	}
}
