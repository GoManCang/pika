package arr

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomInt() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println("rand = ", rand.Intn(100))
	}

	//append
	fmt.Println("----------------")

	// 不声明大小的数组可进行追加元素
	var a []int = []int{1, 2, 3}

	fmt.Println(a)

	a = append(a, 223, 122)
	fmt.Println(a)

	fmt.Println(a[:2])

	//	定义切片时候，指定长度
	var s []int = make([]int, 2)
	//向后追加，向前定义
	s = append(s, 12, 3, 4, 5, 6, 67)
	fmt.Println(s)

	var abc [3]int = [3]int{1, 2, 3}

	fmt.Println(len(abc))
	fmt.Println(cap(abc))

}
