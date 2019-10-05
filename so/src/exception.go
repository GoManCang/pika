package main

import (
	"errors"
	"fmt"
)

func div(a int, b int) (value int, err error) {

	if b == 0 {
		err = errors.New("除数不能为0")
		return 0, err
	}
	return a / b, nil
}

func outofrange() {
	panic("outofrange")
}

func ExceptionTest() {
	// error 返回的是一般性的错误，
	// panic 函数返回的是让程序崩溃的错误，比如数组越界，空指针引用
	v, err := div(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	//panic 异常发生时候，程序会中断运行，随后输出崩溃日志堆栈信息
	//直接调用内置的panic函数也会引发panic异常，接手任何值作为参数
	outofrange()
}
