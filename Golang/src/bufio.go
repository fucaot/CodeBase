package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// 获取一个io对象
	input := bufio.NewScanner(os.Stdin)
	// 开始访问键盘输入
	for input.Scan() {
		// 获取一行键盘输入内容
		counts[input.Text()]++
		fmt.Println(counts)
	}
}
