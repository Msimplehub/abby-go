package main

import (
	"fmt"
	"strconv"
	"time"
)

func goRoutineTest() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go goRoutineTest() // 开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("hello, golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
