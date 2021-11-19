package main

import "fmt"

func main() {
	// 声明一个管道
	var intChan chan int
	// 为chan分配内存，可以存放三个int类型的管道
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值是%v   intChan本身的地址是 %p \n", intChan, &intChan)
	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	// 查看管道长度和容量
	fmt.Printf("管道的长度是%v,容量是%v \n", len(intChan), cap(intChan))
	// 从管道读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2的值是%d", num2)
	// 发现管道的长度减 1
	fmt.Printf("管道的长度是%v,容量是%v ", len(intChan), cap(intChan))
	var num3 = <-intChan
	var num4 = <-intChan
	fmt.Println(num3, num4)
	fmt.Printf("管道的长度是%v,容量是%v ", len(intChan), cap(intChan))
}
