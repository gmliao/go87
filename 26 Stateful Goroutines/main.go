package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	26_Stateful Goroutines
	也可利用兩個channel來達到同步的效果
	一個讀取,一個寫入
*/

// 用來讀取的channel
type ReadOp struct {
	key  int
	resp chan int
}

// 用來寫入的channel
type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	wg := sync.WaitGroup{}

	//用來存放資料的map
	state := make(map[int]int)

	//讀取用的channel
	reads := make(chan ReadOp)

	//寫入用的channel
	writes := make(chan WriteOp)

	readCount := uint64(0)
	writeCount := uint64(0)

	//開啟一個goroutine來處理讀取和寫入的channel
	go func() {
		for {
			select {
			//收到讀取的channel
			case read := <-reads:
				//如果key不存在,回傳變數的預設值
				read.resp <- state[read.key]
				atomic.AddUint64(&readCount, 1)
			//收到寫入的channel
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
				atomic.AddUint64(&writeCount, 1)
			}
		}
	}()

	//開啟100個goroutine來讀取資料
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			time.Sleep(time.Millisecond * 10)
			defer wg.Done()
			read := ReadOp{key: key, resp: make(chan int)}
			reads <- read
			fmt.Printf("Read %d => %d\n", read.key, <-read.resp)
		}(i)
	}

	//開啟100個goroutine來讀取資料
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			write := WriteOp{key: key, val: key, resp: make(chan bool)}
			writes <- write
			<-write.resp
			fmt.Printf("Write %d => %d\n", write.key, write.val)
		}(i)
	}

	wg.Wait()

	//因為readCount和writeCount在goroutine中AddUint64增加,讀取出來就要用LoadUint64讀取
	fmt.Println("readCount:", atomic.LoadUint64(&readCount))
	fmt.Println("writeCount:", atomic.LoadUint64(&writeCount))
}
