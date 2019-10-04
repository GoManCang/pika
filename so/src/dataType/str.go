package dataType

import (
	"fmt"
	"strconv"
	"strings"
)

func swap(a int, b int) {
	a, b = b, a
}

func swapTs(a *int, b *int) {

	*a, *b = *b, *a

}

func sliceJoin() {
	slice := []string{"1234", "2233", "33"}
	// join 数组
	str := strings.Join(slice, ",")
	fmt.Println(str)
	// 查询第一个指定字符的索引位置
	index := strings.Index(str, "33")
	fmt.Println(index)

	// 乘法 字符串
	repeat := strings.Repeat(str, 2)
	fmt.Println(repeat)
	//替换
	replace := strings.Replace(str, "33", "kk", -1)
	fmt.Println(replace)

	//split
	a := "130-430-199"
	split := strings.Split(a, "-")
	fmt.Println(split)

	//trim
	b := "   a c o    "
	trim := strings.Trim(b, " ")
	fmt.Println(trim, len(trim))

	//去掉空字符 转换为数组
	c := "    a u ok ?   "
	s := strings.Fields(c)
	fmt.Println(s)
	fmt.Println(len(s))

}

func TypeChange() {
	str := "pike ping"

	s := []byte(str)
	fmt.Println(s)

	fmt.Println(string(s))

	//转换为字符串
	fmt.Println(strconv.FormatBool(true))

	//转换为bool类型
	b, e := strconv.ParseBool("true")
	if e != nil {
		fmt.Println("类型转换出错")
	} else {
		fmt.Println(b)
	}
	//将其他类型转成字符串添加到字符切片里面
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, false)
	fmt.Println(slice)
	fmt.Println(string(slice))
}

func StringTest() {
	//行参无法改变地址对应的数值
	a, b := 10, 20
	swap(a, b)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", &a)

	swapTs(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

	aa, bb := "abc", "abc_"
	contains := strings.Contains(aa, bb)
	fmt.Println(contains)

	i := strings.Contains(bb, aa)
	fmt.Println(i)

	sliceJoin()

	TypeChange()

}
