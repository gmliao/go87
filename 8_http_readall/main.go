package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	//透過內建的http.Get()方法，取得網頁內容
	resp, err := http.Get("https://www.myip.com.tw/")

	if err != nil {
		panic(err)
	}

	//從resp讀取body
	//body是一個byte slice，所以要用io.ReadAll()轉成string
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	//取得status
	fmt.Println(resp.Status)
}
