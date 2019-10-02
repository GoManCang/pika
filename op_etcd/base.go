package main

import (
	"fmt"
	"math"
)

func main01() {
	/**
	定义和运算
	*/

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

	//	多重复制
	aa, bb, cc := 1, 2.3, "good"
	fmt.Println(aa, bb, cc)

	//	存在自定义变量在多重赋值的情况下是可以进行自动推倒运算的
	var sun int = 10
	var ee int = 23
	dd, ee, sun := 1, 2, 3
	fmt.Println(sun)
	fmt.Println(dd)
	fmt.Println(ee)
	//匿名变量，不接受数据
	_, _ = 10, 20

	fmt.Printf("%d %f", aa, bb)
	fmt.Printf("%d %.2f", aa, bb)
	//输入格式
	var zz int
	fmt.Println(&zz)
	_, _ = fmt.Scan(&zz)

	fmt.Println(zz)

	var a1, a2 int
	a1 = 10
	a2 = 12
	fmt.Println(a1, a2)

}

func main() {

	var a, b int = 1000, 200
	fmt.Println(a / (b * 1.0))

	/***
	变量命名
	1、不允许使用go的关键词
	2、不允许数字开头
	3、允许使用数字 下划线，字母
	4、区分大小写
	*/

	//	浮点类型
	// 字符类型
	// 字符串类型
	//常量使用

	const aa int = 100

	fmt.Println(aa)

	fmt.Println(aa * 9)
	//不允许修改，一般常量是全部大写
	//aa = 200
	//可省略类型，不用类型推倒
	const bb = 1000

	fmt.Println(bb)

	fmt.Println(aa)
	//./base.go:94:14: cannot take the address of aa
	//fmt.Println(&aa)
	//	常量的地址无法访问
	//常量必须定义
	const PI = 3.1415926

	var r float64 = 2

	s := PI * math.Pow(r, 2)
	fmt.Printf("%0.2f", s)

}
