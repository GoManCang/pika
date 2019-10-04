package arr

import "fmt"

func ArrayTest() {
	//声明数组，默认为0
	var a [10]int
	fmt.Println(a)
	a[0] = 100
	fmt.Println(a[0])
	fmt.Println(a[0:3])
	fmt.Println(a[:1])

	var b [10]string = [10]string{"12344455", "12"}
	var bb = [10]string{"12344455", "12"}
	c := [10]string{1: "12344455", 9: "12"}
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(bb, len(bb))

	for i, v := range c {
		fmt.Println(i, v)
	}

	var abc [4]int = [...]int{1, 2, 3, 4}
	var abcd = [...]int{1, 2, 3, 4}
	fmt.Println(abc)
	fmt.Println(abcd)

	i := 0
	j := len(abc) - 1
	fmt.Println(abc)
	fmt.Println(&abc)
	fmt.Println(&abc[0])
	fmt.Println(&abc[1])
	fmt.Println(&abc[2])
	fmt.Println(&abc[3])
	for {
		if i >= j {
			break
		}
		abc[i], abc[j] = abc[j], abc[i]
		i++
		j--
	}

	fmt.Println(abc)

	fmt.Println(&abc)
	fmt.Println(&abc[0])
	fmt.Println(&abc[1])
	fmt.Println(&abc[2])
	fmt.Println(&abc[3])

}
