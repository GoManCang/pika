package main

import "fmt"
import "./easy"

func main() {
	x := fool()
	if x != nil {
		fmt.Println(x)
	}
	fmt.Println("hello world")

	c := easy.Add(1, 2)
	fmt.Println(c)
}

func fool() interface{} {
	return 1
}
