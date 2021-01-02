package main // 要求这里报名必须是 main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int, 1) // 修改 2 为 1 就报错，修改 2 为 3 可以正常运行

	c <- 1

	go func() {
		fmt.Println("in goroutine", <-c)

		fmt.Println("CPU core count : ",runtime.NumCPU())

		fmt.Println("goroutine count : ",runtime.NumGoroutine())
	}()

	c <- 2

	fmt.Println(<-c)

	fmt.Println("main goroutine count : ",runtime.NumGoroutine())
}

// 修改为 1 报如下的错误:
// fatal error: all goroutines are asleep - deadlock!
