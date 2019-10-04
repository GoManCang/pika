package dataType

import (
	"fmt"
	"strconv"
)

//结构体定义在函数外部 外部类的概念？
type User struct {
	id      int
	name    string
	age     int
	address string
	sex     byte
}

func StructTest() {
	type Student struct {
		id   int
		desc string
	}

	//var s Student = Student{
	//	id:   0,
	//	desc: "",
	//}
	//s.id = 1
	//s.desc = "this is test"

	result := make([]Student, 20)

	for i := 0; i <= 10; i++ {
		result[i] = Student{
			id:   i,
			desc: "test " + strconv.FormatInt(int64(i), 10),
		}
	}

	//fmt.Println(result)

	rs := []Student{
		{
			id:   0,
			desc: "",
		},
	}
	rs = append(rs, Student{
		id:   1,
		desc: "this is test",
	})
	fmt.Println(rs)

	var ks []Student

	students := append(ks, Student{
		id:   12,
		desc: "desc",
	})

	ks = append(ks, Student{
		id:   100,
		desc: "ks",
	})

	fmt.Println(students)
	fmt.Println(ks)

}
