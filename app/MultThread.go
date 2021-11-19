package main

import (
	"fmt"
)

// Write Data
func write(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("Write Data", i)
	}
	close(intChan)
}

// Read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Println("读到数据", v)
	}

	// 读取完数据后，即任务完成
	// exitChan 是一个标志，当读完数据后，向exit中存放一个true，主线程除非取出这个true，否则不停止
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go write(intChan)

	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
