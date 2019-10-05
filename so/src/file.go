package main

import (
	"fmt"
	"io"
	"os"
)

func FileTest() {
	//fileName := "c://this.txt"
	/*
		open c://this.txt: no such file or directory
		文件关闭异常： invalid argument
	*/
	fileName := "this.txt"
	fp, err := os.Create(fileName)
	if err != nil {
		//文件创建失败，路径不存在，文件全县，程序打开文件上限
		fmt.Println(err)
	}
	for i := 1; i <= 10; i++ {
		_, _ = fp.WriteString("hello world\n")

	}
	//err = fp.Close()
	//if err != nil {
	//	fmt.Println("文件关闭异常：", err)
	//}

	defer fp.Close()

	//追加
	//count, _ := fp.Seek(0, io.SeekEnd)
	//
	//_, _ = fp.WriteAt([]byte("lishulong"), count)

	//读取内容
	fp, err = os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	//buf := make([]byte, 1024)
	//n, _ := fp.Read(buf)
	//fmt.Println(string(buf[:n]))
	//defer fp.Close()

	//分段读取
	buf := make([]byte, 10)
	for {
		n, err := fp.Read(buf)
		//io.EOF 表示的是文件的结尾
		if err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}

	_ = os.Remove(fileName)

}
