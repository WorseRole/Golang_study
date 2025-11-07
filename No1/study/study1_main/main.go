package main

import (
	"Golang_study/No1/study/study1"
	"fmt"
	"sync"
)

func main() {

	// 指针 值传递 引用传递
	/*
		// 示例1：基本用法
		var num int = 5
		fmt.Printf("增加前的值: %d\n", num)

		study1.ReferencePassing(&num) // 传递变量的地址
		fmt.Printf("增加后的值: %d\n", num)

		// 示例2：使用 new 函数创建指针
		ptr := new(int)
		*ptr = 20
		fmt.Printf("\n使用 new 创建指针:\n")
		fmt.Printf("增加前的值: %d\n", *ptr)

		study1.ReferencePassing(ptr) // 直接传递指针
		fmt.Printf("增加后的值: %d\n", *ptr)

		// 示例3：演示值传递与引用传递的区别
		fmt.Printf("\n值传递 vs 引用传递对比:\n")

		originalValue := 15
		// 值传递（不会修改原始值）
		fmt.Printf("原始值: %d\n", originalValue)
		study1.ValuePassing(originalValue)
		fmt.Printf("值传递后的原始值: %d (未被修改)\n", originalValue)

		// 引用传递（会修改原始值）
		fmt.Printf("原始值: %d\n", originalValue)
		study1.ReferencePassing(&originalValue)
		fmt.Printf("引用传递后的原始值: %d (已被修改)\n", originalValue)
	*/

	// 指针 引用传递 -> slice 切片
	/*
		var data []int = []int{1, 2, 3}
		sliceTemp := &data
		fmt.Println("修改前:", data)
		study1.ReferenceSlice(sliceTemp)
		fmt.Println("修改后:", data)
	*/

	// 协程 两个协程 1-10 一个打印奇数，一个打印偶数

	// 这个就是需要sleep等待 有点傻
	/*
		go func() {
			study1.OddNumber()
		}()

		go func() {
			study1.Even()
		}()
		// 添加等待，让协程有时间执行
		time.Sleep(2 * time.Second)
	*/
	// 这个可以在方法中使用waitGroup 同步
	/**
	go func() {
		study1.OddEvenwithWaitGroup()
	}()
	time.Sleep(1000 * time.Millisecond)
	**/

	// 看看为什么会出现闭包捕获问题！
	// ok 排查完后，原来在1.22以及之后这种写法已经OK
	// 1.22以及之后的版本 每次循环迭代都会创建一个新的循环变量副本
	// var wg sync.WaitGroup
	// for i := range 10 {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}()
	// }
	// wg.Wait()
	//
	// var wg sync.WaitGroup
	// for i := range 10 {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// wg.Wait()

	// tasks := []study1.Task{
	// 	func() {
	// 		// time.Sleep(2 * time.Second)
	// 		fmt.Println("任务 A 完成")
	// 	},
	// 	func() {
	// 		// time.Sleep(1 * time.Second)
	// 		fmt.Println("任务 B 完成")
	// 	},
	// 	func() {
	// 		// time.Sleep(3 * time.Second)
	// 		fmt.Println("任务 C 完成")
	// 	},
	// }
	// study1.RunTasks(tasks)

	// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	// 创建两个实例
	// rect := study1.Rectangle{Width: 4, Height: 3, Name: "rectangle"}
	// circle := study1.Circle{Redius: 3, Name: "circle"}
	// // 声明一个Shape 类型的切片，实现多态
	// shapes := []study1.Shape{rect, circle}
	// for _, shape := range shapes {
	// 	fmt.Printf("形状: %s \n", shape.PrintName())
	// 	fmt.Printf("面积: %.2f \n", shape.Area())
	// 	fmt.Printf("周长: %.2f \n", shape.Perimeter())
	// }

	// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	// person := study1.Person{Name: "leoYan", Age: 27}
	// emp := study1.Employee{Person: person, EmployeeId: 1}
	// emp.PrintInfo()

	// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	// 创建一个不带缓冲的channel
	// ch := make(chan int)
	// // 启动发送 goroutine
	// go study1.SendOnly(ch)
	// // 启动接收goroutine
	// timeout := time.After(10 * time.Second)
	// for {
	// 	select {
	// 	case v, ok := <-ch:
	// 		if !ok {
	// 			fmt.Println("Channel已关闭")
	// 			return
	// 		}
	// 		fmt.Printf("主groutine接收到: %d \n", v)
	// 	case <-timeout:
	// 		fmt.Println("操作超时")
	// 		return
	// 	default:
	// 		fmt.Println("没有数据，等待中...")
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }
	// study1.SendAndReceiveWithWG()

	// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// ch := make(chan int, 20)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go study1.Producer(ch, 100, &wg)
	// go study1.Consumer(ch, &wg)
	// wg.Wait()
	// fmt.Println("===所有任务完成===")

	// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// safeCount := study1.SafeCount{}
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		defer wg.Done()
	// 		for j := 0; j < 1000; j++ {
	// 			safeCount.Increment()
	// 		}
	// 		fmt.Printf(" 协程 %d 完成 %d 次递增 \n", id, 1000)
	// 	}(i)
	// }
	// wg.Wait()
	// actual := safeCount.GetValue()
	// fmt.Printf("结果统计actual: %d \n", actual)
	// 这个也可以
	// study1.IncrementCount()

	// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var wg sync.WaitGroup
	count := study1.AtomicCount{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				count.IncrementAtomicCount()
			}
		}()
		fmt.Printf("i=%d时, count=%d \n", i, count.GetValue())
	}
	wg.Wait()
	fmt.Printf("结束后 count=%d \n", count.GetValue())
}
