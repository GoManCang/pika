package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {

	/*
		go get -u github.com/dgrijalva/jwt-go

	*/

	// 中文不支持 空格也不支持
	a, _ := jwt.DecodeSegment("good s")
	fmt.Println(a)

	fmt.Println(jwt.EncodeSegment(a))

}
