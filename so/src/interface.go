package main

import "fmt"

type HMan interface {
	SayHi()
}

type PMan struct {
	name string
}

func (p *PMan) SayHi() {
	fmt.Println(p.name)
}

func SayHi(h HMan) {
	h.SayHi()
}

func TestInterface() {
	//接口是一种数据类型，可以满足对象信息
	var h HMan
	p := PMan{name: "lishulong"}
	//将对象赋值给接口 必须满足对象实现接口中的方法
	h = &p
	fmt.Println(p)
	h.SayHi()

	// 多态
	SayHi(&p)

	// 接口类型可以接收任意类型的数据
	var i interface{}

	i = 10
	fmt.Println(i)

	i = 3.123
	fmt.Println(i)
	i = "watch"
	fmt.Println(i)
}
