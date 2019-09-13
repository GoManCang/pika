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

	d := easy.DataType(1, 2)
	fmt.Println("abc：", d)

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

func fool() interface{} {
	return 1
}
