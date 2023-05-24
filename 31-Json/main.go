package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	31 Json
	介紹如何使用Json套件,這個套件可以將資料轉換成Json格式,也可以將Json格式轉換成資料
*/

// 定義一個使用者結構,大寫開頭的欄位才會被轉換成Json格式
type User struct {
	Account string `json:"account"` //轉換成Json格式後的欄位名稱
	Name    string `json:"name"`    //轉換成Json格式後的欄位名稱
	Active  bool   `json:"active"`  //轉換成Json格式後的欄位名稱
	Age     int    `json:"age"`     //轉換成Json格式後的欄位名稱
}

func main() {
	tim := User{
		Account: "tim",
		Name:    "Tim",
		Active:  true,
		Age:     18,
	}

	res, err := json.Marshal(tim) //將使用者結構轉換成Json格式
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res)) //印出Json格式的字串

	//將Json格式的字串轉換成使用者結構
	jsonStr := `{"account":"jimmy","name":"Jimmy","active":true,"age":28}`
	jimmy := User{}
	err = json.Unmarshal([]byte(jsonStr), &jimmy)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jimmy)

	// 利用標準輸入將Json格式的字串轉換成使用者結構
	// {"account":"john"}
	var john User
	err = json.NewDecoder(os.Stdin).Decode(&john)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", john)

	// 利用標準輸出將使用者結構轉換成Json格式字串
	json.NewEncoder(os.Stdout).Encode(tim)

}
