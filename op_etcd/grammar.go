package main

import (
	"fmt"
)

func main() {

	var sum int = 0
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	//	7 敲桌子
	//7的倍数 个位是7 十位是7 需要敲桌子

	for i := 1; i <= 100; i++ {
		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println(i, "敲桌子")
		}
	}

	//	水仙花数

	for i := 100; i <= 999; i++ {
		// 百位
		a := i / 100
		b := i / 10 % 10
		c := i % 10

		if i == a*a*a+b*b*b+c*c*c {
			fmt.Println(i)
		}
	}
	count := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			count++

			fmt.Print(j, "*", i, "=", i*j, "\t")
			if i == j {
				//fmt.Println()
				break
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
	// 时间模拟
	//for i := 0; i < 24; i++ {
	//	for j := 0; j < 60; j++ {
	//		for k := 0; k < 60; k++ {
	//			time.Sleep(time.Second)
	//			fmt.Println(i, j, k)
	//		}
	//	}
	//}
	x := 0
	for {
		fmt.Println(x)
		if x > 5 {
			break
		}
		x++
	}

	fmt.Println("hello 1")
	goto SET
	//中间会skip
	fmt.Println("hello 2")
	fmt.Println("hello 3")
SET:
	fmt.Println("hhhh")
	fmt.Println("hhhh")

}
