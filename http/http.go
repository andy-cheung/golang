package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "404")
		return
	}
	fmt.Fprintf(w, "%s!\n", contents)
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
	// 你好
	// test hBttpclient
	fmt.Println("test")
	resp, err := http.Get("http://www.163.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
