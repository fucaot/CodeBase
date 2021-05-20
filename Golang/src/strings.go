package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

// function string.Join splice string
// 拼接字符串
func join() {
	str1 := "abc"
	str2 := "ced"
	str3 := strings.Join(str1, str2)
	fmt.Println(str3)
}
