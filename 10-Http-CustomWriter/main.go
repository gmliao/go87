package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWritter struct{}

// 實作io.Writer的Write()方法
func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {

	//透過內建的http.Get()方法，取得網頁內容
	resp, err := http.Get("https://www.myip.com.tw/")

	if err != nil {
		panic(err)
	}

	//透過io.Copy()方法，將resp.Body的內容寫入到logWritter中
	lw := logWritter{}
	_, err = io.Copy(lw, resp.Body)

	if err != nil {
		panic(err)
	}
}
