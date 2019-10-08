package main

import "fmt"

func main() {
	one()
	fmt.Println('g')
	fmt.Println("g")

}

func fool() interface{} {
	return 1
}

func one() {
	/***
	引用其他包的方法
	add声明一些数据类型的方式
	***/

	var bbcs int = 1234
	fmt.Println("hello ", bbcs)
	x := fool()
	if x != nil {
		fmt.Println(x)
	}
	fmt.Println("hello world")

	//c := Add(2, 3)
	//fmt.Println(c)

	//d := easy.DataType(1, 2)
	//fmt.Println("abc：", d)

	var a, b int
	a = 1
	b = 2
	fmt.Println(a, b)

	fmt.Println("------------------------------------")
	//
	var ad *int
	fmt.Println(ad)

	var ab []int
	fmt.Println(ab)

	var ac map[string]int
	fmt.Println(ac)

	var adc chan int
	fmt.Println(adc)

	var bb func(string) int
	fmt.Println(bb)

	var ccc error
	fmt.Println(ccc)
}
