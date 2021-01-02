package main // 要求这里报名必须是 main

import "fmt"

func main() {
	c := make(chan int, 2) // 修改 2 为 1 就报错，修改 2 为 3 可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
// 修改为 1 报如下的错误:
// fatal error: all goroutines are asleep - deadlock!
