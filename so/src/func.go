package main

import (
	"fmt"
	"strings"
)

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

	//defer fmt.Println("hello deffer 1")
	//defer fmt.Println("hello deffer 2")
	//fmt.Println("after defer")
	/*
		after defer
		hello deffer 2
		hello deffer 1
	*/

	// 函数内部调用多次
	var f1 = func(x, y int) int {
		return x + y
	}

	i2 := f1(1, 20)
	fmt.Println(i2)

	//只调用一次
	a := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(a)

	fmt.Println(strings.Repeat("*", 20))

	// 闭包 使用f1 调用f2
	//fmt.Printf("%T\n", f3(f2, 1, 2))
	i3 := f3(f2, 1, 2)
	//i3()
	f11(i3)

	//f1(i3)

}

func f11(a func()) {
	fmt.Println("f1")
	a()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	fmt.Println("f3")
	return func() {
		f(x, y)
	}
}
