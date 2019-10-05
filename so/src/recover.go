package main

import "fmt"

func t1(i int) {
	var arr [10]int
	// 错误拦截需要定义在出现错误之前
	defer func() {
		i2 := recover()
		fmt.Println(i2)
	}()
	arr[i] = 123
	fmt.Println(arr)
}
func RecoverTest() {

	t1(11)
	fmt.Println("this is test")

}
