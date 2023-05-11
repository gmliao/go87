package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPResponse struct {
	Success bool   `json:"success"`
	IP      string `json:"ip"`
	Type    string `json:"type"`
}

func main() {
	//透過內建的http.Get()方法，取得網頁內容
	resp, err := http.Get("https://api.my-ip.io/ip.json")

	if err != nil {
		panic(err)
	}

	//轉換body到json格式
	var result IPResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("success:%v\n", result.Success)
	fmt.Printf("type:%v\n", result.IP)
	fmt.Printf("myip:%v\n", result.Type)
}
