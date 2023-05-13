package main

import "fmt"

//ch是一個
func sendData(ch chan<- int) {
	//寫入資料後結束
	ch <- 10
}

func main() {

	/*
		var ch1 chan int       // ch1是一个正常的channel，是双向的
		var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
		var ch3 <-chan int     // ch3是单向channel，只用于读int数据
	*/

	//定義一個單向channel send only
	ch := make(chan<- int)

	go sendData(ch)
	fmt.Println("Data sent successfully!")
}
