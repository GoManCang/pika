package main

import "fmt"

func main() {

	a := 10
	b := 20

	b, a = a, b

	fmt.Println(a)
	fmt.Println(b)

}
