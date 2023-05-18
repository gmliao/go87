package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Atomic Counters
	在多個goroutine中,使用atomic操作來增加一個計數器
	使用sync/atomic套件中的AddUint64來增加一個無符號的整數
	使用atomic.LoadUint64來取得最新的計數器值
	如果要減少計數器,使用atomic.AddUint64(&ops, -1)
	如果不使用atomic,則會有race condition的問題
	用以下指令可以檢查race condition
	go run -race main.go

	要使用race,先安裝VS Code的C++插件,https://www.golinuxcloud.com/cgo-tutorial/
	在安裝gcc, https://www.msys2.org/
*/

func main() {

	counter := uint64(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)

		//每個goroutine都會增加計數器2000000次
		go func() {
			for c := 0; c < 2000000; c++ {
				//直接增加計數器,會有race condition的問題
				//counter++
				//使用atomic來增加一個計數器
				atomic.AddUint64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	//正確的計數器值應該是 50*2000000=100000000
	fmt.Printf("Counter: %d\n", counter)
}
