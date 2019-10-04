package dataType

import (
	"fmt"
	"strings"
)

func MapTest() {
	var m map[string]string
	fmt.Println(m)
	//m["name"] = "lishulong"
	//m["age"] = strconv.FormatInt(21, 63)
	//fmt.Println(m)

	//	panic: assignment to entry in nil map

	var k = map[string]string{
		"name": "lishulong",
		"age":  "12",
	}
	fmt.Println(k)

	fmt.Println(k["age"])

	for kk, v := range k {
		fmt.Println(kk, v)
	}

	mm := make(map[string]string)

	fmt.Println("len", len(mm))
	fmt.Println(mm)
	mm["name"] = "lishulong"
	fmt.Printf("%p\n", mm)

	fmt.Println(mm)

	//	map 的长度是自动扩容的
	fmt.Println("cap", len(mm))

	abc := "hello world !!!"
	// 统计字符个数
	// 方式1
	cs := strings.Fields(abc)
	fmt.Println(cs)
	cs = strings.Split(abc, "")
	fmt.Println(cs)
	d_count := make(map[string]int)
	fmt.Println(d_count["1"])
	for _, v := range cs {
		d_count[v] += 1
	}
	for k, v := range d_count {
		fmt.Println(k, v)
	}

	// 方式2 统计个数
	slice := []byte(abc)
	fmt.Println(slice)
	result := make(map[byte]int)
	for _, vs := range slice {
		result[vs] += 1
	}

	for k, v := range result {
		fmt.Printf("%c:%d", k, v)
	}

}
