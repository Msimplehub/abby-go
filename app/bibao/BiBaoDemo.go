package main

import "fmt"

func main() {
	a()
	c := a()
	c()
	c()
	c()

	a()
}

func a() func() int {

	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}

	return b
}
