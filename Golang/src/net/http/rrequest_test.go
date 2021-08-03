package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_ParseForm(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 返回Form表单中指定key的slice
		// 先通过ParseForm解析, 然后访问表单字段
		r.ParseForm()
		fmt.Println(r.Form)
	})
}

func Test_FormValue(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// FormValue 返回Form表单中指定key字段的第一个value
		name := r.FormValue("name")
		// 与FormValue一致, 但仅支持PostForm
		postname := r.PostFormValue("name")
		fmt.Println(name)
		fmt.Println(postname)
	})
}

func Test_ParseMultipartForm(t *testing.T) {
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

func Test_MultipartReader(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 如果是multipart/form-data 或 mutipart混合的请求，则返回multipart.Reader, 否则返回nil
		// 此函数用于代替MutipartForm 将请求作为stream进行处理
		r.MultipartReader()
	})
}
