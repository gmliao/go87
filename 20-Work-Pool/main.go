package main

import (
	"fmt"
	"time"
)

// 工作函式, job 只讀, result 只寫
func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		//模擬處理一個job需要一秒
		time.Sleep(time.Second)
		//結果寫入result channel
		result <- j
		//輸出工作處理結果
		fmt.Println("worker", id, "process ", j)
	}
}

func main() {

	//指派5個工作
	numOfjob := 5

	//建立3個worker
	numOfWorker := 3

	//計時開始
	timeStart := time.Now()

	//建立job channel用來指派工作,必須用buffer channel,因為指派工作的時候還沒有接收者
	job := make(chan int, numOfjob)

	//建立result channel用來接收結果
	result := make(chan int, numOfjob)

	//建立3個worker
	for i := 0; i < numOfWorker; i++ {
		go worker(i, job, result)
	}

	//指派5個工作
	for j := 0; j < numOfjob; j++ {
		job <- j
	}

	//讀取出result channel的結果,必須用for迴圈,因為result channel沒有close, for range只能用在close的channel
	for r := 0; r < numOfjob; r++ {
		fmt.Println("get result:", <-result)
	}

	fmt.Println("time:", time.Since(timeStart))
}
