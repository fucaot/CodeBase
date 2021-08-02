package http

import (
	"net/http"
	"testing"
)

func Test_Write(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 使用Write写入byte[]切片
		// 如果未设置Content-Type, 则前512字节用来检测Content-Type
	})
}

func NotFound(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
}
