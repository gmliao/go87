package main

import (
	"fmt"
	"time"
)

/*
	Rate Limiting 1
	假設要限制處理request的頻率,limiter 每兩百毫秒會放行一個request
	要處理reqeust的時候先讀取limiter,如果limiter有值就放行,沒有值就等待
	這樣就可以限制處理request的數量,避免過多的request造成系統負擔
*/

func handler(id int) {
	fmt.Printf("request id %d is done.\n", id)
}

func main() {

	jobNum := 5

	request := make(chan int, 5)

	//新增jobNum個job
	for i := 0; i < jobNum; i++ {
		request <- i
	}

	//必須request,下方for range才知道channel已經關閉,否則會一直等待
	close(request)

	//初始化limiter,兩百毫秒觸發依次
	limiter := time.Tick(time.Millisecond * 200)

	for r := range request {
		fmt.Printf("request id %d is waiting.\n", r)
		<-limiter
		handler(r)
	}
}
