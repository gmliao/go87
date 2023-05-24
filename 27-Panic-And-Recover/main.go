package main

import (
	"fmt"
)

/*
	27 Panic And Recover
	發生錯誤的時候可以用panic來中斷程式
	但是在goroutine中,如果發生panic,整個程式就會中斷
	所以可以用recover來捕捉panic,並且讓程式繼續執行
	可以用在http server中,讓程式不會因為一個請求的錯誤而中斷
*/

func request(error bool) {
	// recover只能在defer中使用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("request fail, Error:", err)
		}
	}()

	// 這裡會發生panic
	if error {
		panic("something wrong")
	} else {
		fmt.Println("request success")
	}
}

func main() {
	request(true)
	request(false)
	request(false)

	fmt.Println("all request success")
}
