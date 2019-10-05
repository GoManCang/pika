package dataType

import (
	"fmt"
)

// 定义函数类型
// 定义结构体名称
// 为已存在的数据类型起别名

type Int int

// 方法
func (a Int) add(b Int) Int {
	return a + b
}

func (u User) getAge() int {
	u.age = 30
	return u.age
}

func (u *User) getAge1() int {
	u.age = 20
	return u.age
}

type Human struct {
	name string
	age  int
}

type RMan struct {
	Human
	hobby string
}

type EMan struct {
	Human
	workTime int
}

func (h *Human) SayHi() {
	fmt.Println(h.name, h.age)
}

func (r *RMan) SayHello() {
	//r.Human.SayHi()
	r.SayHi()
	fmt.Println(r.hobby)
}

func TypeTest() {
	// 类型别名会在编译时进行转换
	// 预处理- 包展开，替换数据类型，去掉注释>
	// -> 编译（代码有错会提示）编译成汇编文件
	// ->汇编，将汇编文件转成二进制文件
	// ->连接 将支持的库连接到程序中，链接为可执行程序
	var a Int = 10
	var b int = 20

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a + Int(b))

	fmt.Println(a.add(Int(b)))

	//
	user := User{
		id:      0,
		name:    "贾宝玉",
		age:     21,
		address: "大观园",
		sex:     0,
	}

	// 值传递
	fmt.Println(user.getAge())
	fmt.Println(user.age)

	//地址传递
	fmt.Println(user.getAge1())
	fmt.Println(user.age)

	r := RMan{
		Human: Human{
			"name",
			12,
		},
		hobby: "羽毛球",
	}

	r.SayHi()
	r.SayHello()

}
