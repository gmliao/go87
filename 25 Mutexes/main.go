package main

import (
	"fmt"
	"sync"
)

/*
	24 Mutexes
	1. Mutexes 是一種互斥鎖,用來保護共享資源不被同時訪問
	2. 一個goroutine在訪問共享資源時,需要先鎖定Mutexes,訪問完畢後再解鎖
	3. Mutexes 只有兩種狀態,鎖定和未鎖定,使用Mutexes.Lock()鎖定,使用Mutexes.Unlock()解鎖
	4. Mutexes 可以在不同的goroutine中進行鎖定和解鎖,但是鎖定和解鎖的次數必須相同
	5. Mutexes 可以在同一個goroutine中進行鎖定和解鎖,但是鎖定和解鎖的次數必須相同

	下面範例是一個簡單的計數器,使用Mutexes保護共享資源
*/

type Container struct {
	lock  sync.Mutex
	datas map[string]int
}

func (c *Container) Increase(k string) {
	// 如果沒上鎖,會造成資料競爭馬上掛掉
	c.lock.Lock()
	defer c.lock.Unlock()
	c.datas[k]++
}

func main() {

	wg := sync.WaitGroup{}

	// lock 不需要初始化也可以
	c := Container{datas: make(map[string]int)}

	updateContainer := func(c *Container, key string) {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			c.Increase(key)
		}
	}

	wg.Add(3)

	go updateContainer(&c, "a")
	go updateContainer(&c, "a")
	go updateContainer(&c, "a")

	wg.Wait()

	fmt.Println(&c.datas)
}
