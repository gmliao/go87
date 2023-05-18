package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	有五個工作要執行，每個工作需要花費的時間不同，利用WaitGroup來等待所有工作完成後才繼續執行下去
*/

func worker(id int, results chan<- int) {
	workTime := time.Duration(rand.Intn(5))
	time.Sleep(workTime * time.Second)
	results <- id
	fmt.Printf("worker %d has done his job, spent %d secs.\n", id, workTime)
}

func main() {

	numOfWorker := 5
	result := make(chan int, numOfWorker)

	wg := sync.WaitGroup{}

	//建立五個工作
	for i := 0; i < numOfWorker; i++ {
		wg.Add(1)
		//用匿名函式來呼叫worker,就不用傳送wg,這樣可以避免傳送wg時,被其他函式誤用
		go func(i int) {
			defer wg.Done()
			worker(i, result)
		}(i)
	}

	//等待所有工作完成,會等到wait.Done()被呼叫五次
	wg.Wait()

	fmt.Println("All jobs are done.")

	//取出結果
	for i := 0; i < numOfWorker; i++ {
		fmt.Printf("result %d, worker %d .\n", i, <-result)
	}
}
