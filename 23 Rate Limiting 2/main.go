package main

import (
	"fmt"
	"time"
)

/*
	Rate Limiting
	假設要限制處理request的頻率,增加一個burstyLimiter,200毫秒增加一個最多三個
*/

func handleRequest(id int) {
	fmt.Printf("request id %d is done.\n", id)
}

func main() {
	//時間型態的channel,最多三個
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//發起一個goroutine,每200毫秒增加一個busterLimiter
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			fmt.Printf("time is up, one bursterLimiter\n")
			burstyLimiter <- t
		}
	}()

	requestNum := 5
	burstyRequest := make(chan int, requestNum)
	//建立5個工作
	for i := 0; i < requestNum; i++ {
		burstyRequest <- i + 1
	}
	close(burstyRequest)

	//等待5個request結束
	for req := range burstyRequest {
		//等待burstyLimiter有東西才會繼續執行
		<-burstyLimiter
		handleRequest(req)
	}
}
