package http

import (
	"net/http"
	"testing"
)

func Test_Write(t *testing.T) {
	http.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		// 使用Write写入byte[]切片
		// 如果未设置Content-Type, 则前512字节用来检测Content-Type
		str := `<?xml version=\"1.0\" encoding=\"UTF-8\"?>
		<root>
			<name >asaa</name >
			<das>wjkkj</das>
		</root>`
		w.Write([]byte(str))

		// 执行 curl -i localhost:8080/write 试试
	})
	http.ListenAndServe("localhost:8080", nil)
}

func TestNotFound(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
}
