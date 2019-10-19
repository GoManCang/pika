package main

import "fmt"

func Test0() int {
	x := 8

	defer func() {
		x++
	}()
	return x

}

func Test01() (a int) {
	a = 8

	defer func() {
		a++
	}()
	return a

}
func Test02() (a int) {

	defer func() {
		a++
	}()
	return 8

}
func Test03() (a int) {
	b := 5
	defer func() {
		b++
	}()
	// a = b
	// b++
	return b

}
func Test04() (a int) {
	defer func(a int) {
		a++
	}(a)
	// a = 5
	// a = 6  a是 a的副本，值传递
	return 5

}

func Test05() (a int) {
	defer func(a *int) {
		*a++
	}(&a)
	// a = 5
	// a = 6  a==a，地址传递
	return 5

}
func main() {

	fmt.Println(Test0())
	fmt.Println(Test01())
	fmt.Println(Test02())
	fmt.Println(Test03())
	fmt.Println(Test04())
	fmt.Println(Test05())
	/*
		 	8
			9
			9
			5
			5
			6

	*/

}
