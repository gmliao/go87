package main

func main() {
	// 這個程式會造成 deadlock
	// 因為 main goroutine 會一直等待 c channel 被讀取
	// 但是沒有其他 goroutine 會去讀取 c channel
	// 所以 main goroutine 會一直等待下去
	// 造成 deadlock
	c := make(chan string)
	c <- "Hi there!"
}
