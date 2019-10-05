package base

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

func Test() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func main() {

	test()
	a, v := add(1, 3)
	fmt.Println(a, v)

	s := test_(1, 2, 3, 4, 5)
	fmt.Println(s)

	//阶乘
	// 1*2*3*4*5
	sum_ := 1
	var vs func(int)
	vs = func(n int) {
		if n == 1 {
			return
		}
		vs(n - 1)
		sum_ *= n
	}

	vs(5)
	fmt.Println(sum_)
	fmt.Println(Test())

	c, a, b := Test()
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)

	//函数定义参数
	type Ftest func() (int, int, int)
	// 声明函数
	var abc Ftest
	//赋值函数
	abc = Test
	fmt.Println(abc())

}
