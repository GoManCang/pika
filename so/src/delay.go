package main

import "fmt"

func delayFunc() {
	fmt.Println("delayFunc...")
}

func DelayTest() {
	/*
		1、defer 语句只能出现在函数的内部，延迟一个函数的执行
		2、应用场景，文件操作
	*/
	defer delayFunc()
	fmt.Println("hello defer")

	a, b := 10, 20
	//不传递参数，外部会影响内部
	f := func() {
		fmt.Println("f:", a+b)
	}
	defer f()

	fc := func(a int, b int) {
		fmt.Println("fc:", a+b)
	}
	defer fc(a, b)
	a = 100
	b = 200
}
