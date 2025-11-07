package study1

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
	PrintName() string
}

// 定义矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
	Name   string
}

// 实现Shape 接口的方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) PrintName() string {
	return r.Name
}

// 定义圆形结构体
type Circle struct {
	Redius float64
	Name   string
}

// 实现Shape接口的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Redius
}

func (c Circle) PrintName() string {
	return c.Name
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

// 定义Person 员工结构体
type Person struct {
	Name string
	Age  int
}

// 定义 Employee 结构体
type Employee struct {
	Person     Person
	EmployeeId int
}

func (emp Employee) PrintInfo() {
	fmt.Printf("员工姓名为：%s, 年龄为: %d, 员工ID为: %d", emp.Person.Name, emp.Person.Age, emp.EmployeeId)
}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 往协程中发送channel 的函数
func SendOnly(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("发送: %d \n", i)
	}
	close(ch)
}

// 接收channel 的函数
func ReceiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d \n", v)
	}
}

// 直接用WaitGroup 等待完成
func SendAndReceiveWithWG() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	// 等待两个协程
	wg.Add(2)

	// 发送协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("发送: %d \n", i)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Printf("接收到: %d \n", v)
		}
	}()

	wg.Wait()
	fmt.Println("所有通信完成")
}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func Producer(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	fmt.Printf("生产者开始发送:%d 个整数 \n", count)
	for i := 1; i <= count; i++ {
		ch <- i
		fmt.Printf("生产: %d (通道中还有%d/%d 个元素)\n", i, len(ch), cap(ch))
		// 假装生产耗时50ms
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("生产者完成所有发送任务")
}

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("消费者开始接收整数")
	receivedCount := 0
	for value := range ch {
		receivedCount++
		fmt.Printf("消费: %d(通道中还剩 %d/%d 个元素)\n", value, len(ch), cap(ch))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("消费者完成，共接收 %d 个整数 \n", receivedCount)
}

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 编写一个 计数器的接口 目前只有递增一个方法
type Counter interface {
	Increment(i int)
}
type SafeCount struct {
	mu    sync.Mutex
	value int
}

func (count *SafeCount) Increment() {
	count.mu.Lock()
	defer count.mu.Unlock()
	count.value++
}
func (count *SafeCount) GetValue() int {
	count.mu.Lock()
	defer count.mu.Unlock()
	return count.value
}

var (
	mu    sync.Mutex
	count int
)

func IncrementCount() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				fmt.Printf("正在自增中: %d\n", count)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("最中count为: %d\n", count)
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
type AtomicCount struct {
	value int64
}

func (count *AtomicCount) IncrementAtomicCount() {
	atomic.AddInt64(&count.value, 1)
}

func (count *AtomicCount) GetValue() int64 {
	return count.value
}
