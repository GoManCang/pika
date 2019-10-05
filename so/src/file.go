package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func createFIle(fileName string) *os.File {
	/*
		//fileName := "c://this.txt"
			open c://this.txt: no such file or directory
			文件关闭异常： invalid argument
	*/

	fp, err := os.Create(fileName)
	if err != nil {
		//文件创建失败，路径不存在，文件全县，程序打开文件上限
		fmt.Println(err)
	}
	for i := 1; i <= 10; i++ {
		_, _ = fp.WriteString("hello world" + strconv.FormatInt(int64(i), 10) + "\n")

	}
	defer fp.Close()
	return fp
}

func readFile(fileName string) {
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	buf := make([]byte, 1024)
	n, _ := fp.Read(buf)
	fmt.Println(string(buf[:n]))
	defer fp.Close()
}

func readFileBySegment(fileName string) {
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	buf := make([]byte, 10)
	for {
		n, err := fp.Read(buf)
		//io.EOF 表示的是文件的结尾
		if err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}
	defer fp.Close()
}

func appendFile(fp *os.File) {
	//追加
	count, _ := fp.Seek(0, io.SeekEnd)
	_, _ = fp.WriteAt([]byte("lishulong"), count)

}

func cacheReadFile(fileName string) {
	fp, err := os.OpenFile(fileName, os.O_RDONLY, 6)
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer fp.Close()

	r := bufio.NewReader(fp)
	//slice, _ := r.ReadBytes('\n')
	//fmt.Println(string(slice))
	//
	//slice, _ = r.ReadBytes('\n')
	//fmt.Println(string(slice))
	for {
		//acsii 的任意值
		slice, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(string(slice))
	}
}

func FileTest() {

	fileName := "this.txt"
	fp := createFIle(fileName)
	err := fp.Close()
	if err != nil {
		fmt.Println("文件关闭异常：", err)
	}

	//追加
	appendFile(fp)
	//读取内容
	readFile(fileName)
	//分段读取
	readFileBySegment(fileName)
	//

	fmt.Println(strings.Repeat("-", 20))

	cacheReadFile(fileName)

	_ = os.Remove(fileName)

}
