package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个map
var (
	myMap = make(map[int]int, 10)
	// 声明一个全局变量互斥锁
	// lock 是一个全局互斥锁
	// sync 是包：sychorinized
	// Mutex : 互斥
	lock sync.Mutex
)

// 定义一个函数计算n！
func factorial(num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	// 将计算出的结果存入map中
	// 访问 myMap 前，加锁
	lock.Lock()
	myMap[num] = res
	// 访问完，释放所
	lock.Unlock()
}
func main() {

	// 开启 200 个协程完成这个任务
	for i := 1; i <= 200; i++ {
		go factorial(i)
	}
	// 休眠 10S
	time.Sleep(time.Second * 10)
	// 输出map中的结果
	lock.Lock()
	for index, value := range myMap {
		fmt.Printf("map[%d]=[%d]\n", index, value)
		//fmt.Println(index,"的阶乘为",value)
	}
	lock.Unlock()
}
