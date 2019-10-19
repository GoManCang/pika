package main

import "fmt"

func main() {

	var a *int
	fmt.Println(a)
	b := 12
	a = &b

	*a = 13
	fmt.Println(a, &a, *a)
	// 申请一个地址
	var i = new(int)
	fmt.Println(i)
	*i = 12
	fmt.Println(i, &i, *i)
}
