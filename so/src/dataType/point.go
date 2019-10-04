package dataType

import "fmt"

func PointTest() {
	//地址就是指针
	//
	a := 10
	fmt.Println(&a)

	p := &a
	fmt.Printf("%T\n", p)

	//0xc0000a0000
	//*int
	var c *int = p
	fmt.Println(c)

	fmt.Println(*p)

	//无效的空内存地址
	//声明指针变量
	var cc *int = nil
	//0x0 内存编号是0～255的空间为系统占用，不允许用户访问
	fmt.Printf("%p\n", cc)
	//panic: runtime error: invalid memory address or nil pointer dereference
	//*cc = 10
	//fmt.Println(*cc)

	//实例化 开辟空间
	cc = new(int)
	*cc = 99
	fmt.Println(cc)
	fmt.Println(*cc)

	//野指针，指针变量指向了一个位置空间，回报错
	//invalid indirect of int(824633819280) (type int)
	//var pp *int = *int(0xc000018090)
	//fmt.Println
	//程序中不允许出现这种野指针
	//允许空指针

	//	数组指针 / 指针数组
	var ac []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var pc *[]int
	pc = &ac
	fmt.Println(pc)
	//(*pc)[1] = 22
	//pc[1] = 22
	fmt.Println(*pc)
	//指针变量和数组存储的类型要一致。声明个数的化也要一致
	abc := 10
	bbc := 20
	cbc := 30
	var arr [3]*int = [...]*int{
		&abc, &bbc, &cbc,
	}
	fmt.Println(arr)
	fmt.Println(*arr[1])

	//结构体和指针
	u := User{
		id:      0,
		name:    "",
		age:     0,
		address: "",
		sex:     0,
	}
	test(&u)
	fmt.Println(u)
}

func test(p *User) {
	p.name = "lishulong"
}
