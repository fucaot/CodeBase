package main

import (
	"bufio"
	"fmt"
	"os"
)

// the os package can help you get the input external input, by the 'os.Args' it's a string slice,
// stone its name in first element and the remaning elements stone the parameters passed when program start
// 你可以通过os包来进行获取外部输入,通过'os.Args', 它是一个string类型的切片，
// os.Args 的第一个元素存储的是其本身的名称, 剩余元素则是存放当程序启动时传递进的参数

// commasd and output/ 输入以及输出:
// command:  os
// command:  os1
// command:  os2
// commands:  os os1 os2
// os.Args[0] name:  /var/folders/s3/3lw7rn712ls2m3zt3smkzrn40000gn/T/go-build1584056910/b001/exe/os
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("command: ", os.Args[i])
	}
	// Println all input commands
	// 打印出所有输入参数
	fmt.Println("commands: ", s)
	fmt.Println("os.Args[0] name: ", os.Args[0])
}
