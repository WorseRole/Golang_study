package study1

import (
	"fmt"
	"sync"
	"time"
)

/*
*
题目 ：编写一个Go程序，定义一个函数，
该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。

考察点 ：指针的使用、值传递与引用传递的区别。
*
*/

// 接收整数指针作为参数，修改指针指向的值
func ReferencePassing(num *int) {
	if num != nil {
		*num += 10
	}
}

// 值传递函数 - 接收整数值的副本
func ValuePassing(num int) {
	num += 10
}

/**
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
**/

func ReferenceSlice(sliceTemp *[]int) {
	if *sliceTemp != nil {
		for i := 0; i < len(*sliceTemp); i++ {
			(*sliceTemp)[i] = (*sliceTemp)[i] * 2
		}
	}
}

/**
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
**/
// 奇数
func OddNumber() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Printf("odd-i:%d \n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// 偶数
func Even() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("even-i:%d \n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func OddEvenwithWaitGroup() {
	var wg sync.WaitGroup
	// 等待两个
	wg.Add(2)

	go func() {
		defer wg.Done()
		OddNumber()
	}()

	go func() {
		defer wg.Done()
		Even()
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有数字打印完成")
}

// Task 是一个任务类型，定义为无参函数
type Task func()

// RunTasks 并发执行所有任务，并统计执行时间
func RunTasks(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for i := 0; i < len(tasks); i++ {
		go func() {
			defer wg.Done()
			start := time.Now()
			tasks[i]()
			duration := time.Since(start)
			fmt.Printf("任务 %d 执行完成，用时 %v \n", i, duration)
		}()
	}

	// for i, task := range tasks {
	// 	go func(u int, t Task) {
	// 		defer wg.Done()

	// 		start := time.Now()
	// 		t()
	// 		duration := time.Since(start)

	// 		fmt.Printf("任务 %d 执行完成，用时 %v \n", i+1, duration)
	// 	}(i, task)
	// }
	wg.Wait()
	fmt.Println(("所有任务执行完成"))
}
