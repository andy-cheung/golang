package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for index := 0; index < 127; index++ {
		url := fmt.Sprintf("http://z.m.com:8501/get_server_addr?platformid=111111111&serverid=%d", index)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		io.Copy(os.Stdout, resp.Body)
		fmt.Print("\n")
		resp.Body.Close()
	}
}
