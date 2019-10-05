package main

import (
	"fmt"
)

type User struct {
	name string
}

func AssertTest() {

	//空接口可存储任意类型的数值
	//可使用类型断言来知道是那个 来进行逻辑处理
	//语法 判断是否是该类型的变量
	//value,ok = element.(T)
	//value：变量数值，
	// ok：bool
	// element：interface 变量
	// T是断言类型
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = User{name: "lishulong"}
	i[2] = "hello"

	fmt.Println(i)

	for index, data := range i {
		if v, ok := data.(int); ok == true {
			fmt.Println(index, v, ok)
		} else if v, ok := data.(User); ok == true {
			fmt.Println(index, v, ok)
		}

	}

}
