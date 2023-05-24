package main

import (
	"fmt"
	"time"
)

func main() {
	//建立一個ticker, 每秒鐘執行一次
	ticker := time.NewTicker(time.Second * 1)
	done := make(chan bool)

	//建立一個goroutine, 每秒鐘執行一次
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done!")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	//執行5秒後停止ticker
	time.Sleep(time.Second * 5)
	ticker.Stop()

	//通知goroutine停止,會等待done channel接收到值,才會繼續執行
	done <- true

	fmt.Println("Ticker stopped")
}
