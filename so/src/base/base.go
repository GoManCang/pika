package base

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

func main2() {

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

func main3() {

	const (
		a = iota
		//v, c = iota, iota
		f
		g
		k = 10
		l
		p = iota
	)

	fmt.Println(a)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(p)

	/*
		0
		1
		2
		10
		10
		5
	*/

}

func main4() {
	a, b := 10.0, 0.0000000000
	s := a / b
	//panic: runtime error: integer divide by zero
	//+Inf 无穷大么？？
	//取余不能对浮点数进行
	//c := a % b
	//fmt.Println(c)
	fmt.Println(s)

}

func main5() {

	a := 10
	a++
	//不能进行赋值运算
	//c := a++
	fmt.Println(a)

	const c = 10

	//c++
	//	常量不允许自增自减

}

func main7() {
	//	类型转换
	// 高类型的转成低类型，可能会丢失静笃
	// 数据溢出，符号发生改变
	var a int = 2568
	fmt.Println(int8(a))
	fmt.Println(int16(a))

	a = 107653 // 几天几时几分
	fmt.Println("秒：", a%60)
	fmt.Println("分：", a/60%60)
	fmt.Println("时：", a/60/60%24)
	fmt.Println("天：", a/60/60/24%365)

}

func main8() {

	a, b := 10, 20

	fmt.Println(a > b && a > 20 || b > 20)
	fmt.Println(a < b && a < 20 || b < 20)

	fmt.Println(a > b || a < 20 && b < 20)

	//	逻辑与 高于逻辑或

}

func main9() {
	var a int = 20
	// 取地址运算符
	fmt.Println(&a)
	// * 取运算符，p 指针变量
	c := &a
	fmt.Println(*c)
	//通过指针间接修改变量的值
	*c = 200
	fmt.Println(a)

	//	（） 括号[] 数组下标，.结构体
	//	位运算 & | ^ ~
	//	位移运算>> <<
	// 通信>-

	const (
		aa        = iota
		bb, cc, d = iota, iota, iota
	)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
}

func main() {
	var s int
	_, _ = fmt.Scan(&s)

	switch s > 10 {
	case false:
		fmt.Println("s < 10")
	case true:
		fmt.Println("s > 10")
	}

	switch {
	case s == 10:
		fmt.Println("s==10")
	case s > 10:
		fmt.Println("s>10")
	default:
		fmt.Println("s<10")
	}

	fmt.Println("---------------------")
	switch s {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough

	case 3, 4, 5, 7:
		fmt.Println("3,4,5,7")
	default:
		fmt.Println(s)

		/**
		1
		s < 10
		s<10
		---------------------
		1
		2
		3,4,5,7

		*/

	}

}
