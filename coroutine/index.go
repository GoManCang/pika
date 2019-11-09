package main

import (
	"fmt"
	"time"
)

func TestCoroutine(a int, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

func Add(a, b int) int {
	return a + b
}

func main() {
	//var c int
	//c = Add(100, 200)
	for i := 0; i < 100; i++ {
		go TestCoroutine(200, i)
	}
	//fmt.Println(c)

	time.Sleep(time.Second)

}
