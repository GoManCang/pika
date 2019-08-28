package main

import "fmt"

func main() {
	x := fool()
	if x != nil {
		fmt.Println(x)
	}
	fmt.Println("hello world")
}

func fool() interface{} {
	return 1
}
