package main

import "fmt"

func Test(a int, b int) (ret int) {
	/*
		参数可以命名 也可以不命名
	*/
	ret = a + b
	return ret
}

func test(a int, b int) int {
	return a + b
}

func test1(a, b, c int) int {
	/*
		类型声明简写
	*/
	return a + b + c
}

func test3(a string, y ...int) {
	/*
		允许传送多个参数 可变参数
	*/
	fmt.Println(a, y)

}

func main() {
	ret := Test(1, 2)
	fmt.Println(ret)

	i := test(12, 33)
	fmt.Println(i)

	test3("1", 2, 3, 4, 5)

	defer fmt.Println("hello deffer 1")
	defer fmt.Println("hello deffer 2")
	fmt.Println("after defer")
	/*
		after defer
		hello deffer 2
		hello deffer 1
	*/

}
