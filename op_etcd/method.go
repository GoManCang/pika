package main

import "fmt"

func test() {
	fmt.Println("this is test")
}

func add(i int, j int) (int, int) {
	return i + j, i * j

}

func test_(arr ...int) int {

	for i, v := range arr {
		fmt.Println(i, v)
	}
	return len(arr)
}

func main() {

	test()
	a, v := add(1, 3)
	fmt.Println(a, v)

	s := test_(1, 2, 3, 4, 5)
	fmt.Println(s)

}
