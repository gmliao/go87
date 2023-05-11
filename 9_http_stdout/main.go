package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	//透過內建的http.Get()方法，取得網頁內容
	resp, err := http.Get("https://www.myip.com.tw/")

	if err != nil {
		panic(err)
	}

	//透過io.Copy()方法將resp.Body的內容輸出到終端機os.Stdout
	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		panic(err)
	}
}
