package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Get(url string) {
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

func Handle() {
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

func FileServer() {
	// 使用http.FileServer方法来构建简单的文件系统
	// 此时访问localhost:8080/index.html即可以访问代码wwwroot目录下的页面
	// 方法参数为基于root的文件系统, 此处使用http.Dir方法构建, 参数为文件系统根目录路径
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}

func ParseForm() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 返回Form表单中指定key的slice
		// 先通过ParseForm解析, 然后访问表单字段
		r.ParseForm()
		fmt.Println(r.Form)
	})
}

func FormValue() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// FormValue 返回Form表单中指定key字段的第一个value
		name := r.FormValue("name")
		// 与FormValue一致, 但仅支持PostForm
		postname := r.PostFormValue("name")
		fmt.Println(name)
		fmt.Println(postname)
	})
}

func ParseMultipartForm() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// multpart/form-data最常见的应用场景就是上传文件
		// 调用ParseMultpartForm进行解析, 参数, 一次访问的子节长度
		r.ParseMultipartForm(1024)

		// 获取下载文件
		upload := r.MultipartForm.File["upload"][0]
		file, err := upload.Open()
		if err != nil {
			fmt.Println(err)
		}
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
		}
		w.Write(fileData)
		// 获取字段
		w.Write([]byte(r.MultipartForm.Value["name"][0]))
	})
}
