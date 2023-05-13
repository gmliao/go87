package main

import "fmt"

//接收ch的資料
func receiveData(ch <-chan int) {
	//等待接收資料
	data := <-ch
	fmt.Println("Data received:", data)
}

func main() {
	//定義一個只能發送int的雙向channel
	ch := make(chan int)

	//開啟一個goroutine接收ch的資料
	go receiveData(ch)

	//主goroutine發送資料到ch
	ch <- 10
	fmt.Println("Data sent successfully!")
}
