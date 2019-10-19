package main

import "fmt"

func main() {
	a := make([]int, 2, 10)
	fmt.Println(len(a), cap(a))

	s1 := make([]int, 0)
	//append

	s1 = append(s1, 1)
	s1 = append(s1, 1)
	s1 = append(s1, 1)
	s1 = append(s1, 1)

	for i := 0; i < 100; i++ {
		s1 = append(s1, i)
	}

	fmt.Println(len(s1), cap(s1))

	// copy
	a = []int{1, 2, 3, 4, 5}
	var c []int = make([]int, 2, 3)

	copy(c, a)
	fmt.Println(a, c)

	a[1] = 1000
	fmt.Println(a, c)

	s2 := []int{1, 2, 3, 4, 5}

	// 切片是一段连续的地址空间，是一个底层数组
	x1 := s2[:]
	//[1,3,4,5] 前四个地址的数值
	x1 = append(s2[:1], s2[2:]...)
	//[1 3 4 5 5] [1 3 4 5]
	fmt.Println(s2, x1)

}
