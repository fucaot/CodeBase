package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_Get(t *testing.T) {
	url := "www.github.com"
	// 通过url 进行httpGet请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("reading err: ", err)
		os.Exit(1)
	}
	status := resp.Status
	fmt.Println(string(body[:]))
	fmt.Println(string(status[:]))
}

func Test_Handle(t *testing.T) {
	// http.Handle方法结构, 此方法默认使用DefaultMux, 因此定义Handler为nil的server对象均可使用
	// func Handle(pattern string, handler handler)
	server := http.Server{
		Addr: "localhost: 8081",
	}
	// 通过http.Handle实现多handle
	http.Handle("/hello", http.HandlerFunc(handlerHello))
	http.Handle("/about", http.HandlerFunc(handlerAbout))
	http.Handle("/process", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	var hh = helloHandler{}
	var ah = aboutHandler{}
	http.Handle("/about", hh)
	http.Handle("/about", ah)
	server.ListenAndServe()
}

type helloHandler struct{}

func (m helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type aboutHandler struct{}

func (m aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}
func handlerHello(w http.ResponseWriter, r *http.Request) {}
func handlerAbout(w http.ResponseWriter, r *http.Request) {}

func Test_FileServer(t *testing.T) {
	// 使用http.FileServer方法来构建简单的文件系统
	// 此时访问localhost:8080/index.html即可以访问代码wwwroot目录下的页面
	// 方法参数为基于root的文件系统, 此处使用http.Dir方法构建, 参数为文件系统根目录路径
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
