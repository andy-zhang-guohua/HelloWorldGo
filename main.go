package main // 要求这里报名必须是 main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 表示让 CPU 把时间片让给别人，下次某个时候继续恢复执行该 goroutine
		fmt.Println(s)
	}
}

func main() {
	go say("world") // 开一个新的 Goroutines 执行
	say("hello") // 当前 Goroutines 执行
}