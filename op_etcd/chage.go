package main

import "fmt"

func main() {

	a := 10
	b := 20

	b, a = a, b

	fmt.Println(a)
	fmt.Println(b)

	var c int
	c = a
	a = b
	b = c
	fmt.Println(a)
	fmt.Println(b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)
}
