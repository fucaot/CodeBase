package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL.Path = %q", r.URL.Path)
}

func get(url string) {
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
