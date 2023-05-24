package main

import (
	"fmt"
)

func main() {

	// 定義一個buffer channel，長度為2
	// 可以允許沒有接收者的情況下發送
	// 但是當channel的長度超過2時，就會發生deadlock
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
