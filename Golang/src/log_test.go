package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

type fllog log.Logger

const filePath = "/log.log"

// output: 无输出
// log.log:
// go-epcc2021/05/20 15:35:13 my first log
func Test_New(t *testing.T) {
	// 此处获取File类型返回值, 实现了io.Write接口, 在获取log对象时需要用到
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("err: ", err)
	}
	// 获取log对象，三个参数分别为:
	// 写入文件的io.Write实现对象
	// 前缀
	// 标识，此处的LStdFlags 代表日期 + 时间
	fllog := log.New(file, "go-epcc", log.LstdFlags)

	// 写入log
	fllog.Println("my first log")
}
