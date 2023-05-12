package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//channel使用時機,goroutine之間需要溝通

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//同步版本會block住，直到所有的link都執行完畢
	for _, link := range links {
		CheckLinkBlocking(link)
	}

	c := make(chan string)

	//for迴圈，每個link都會執行一次CheckLink()
	for _, link := range links {
		//建立goroutine，並且執行CheckLink()，並且傳入link參數
		go CheckLink(link, c)
	}

	counter := 0
	//等待channel回傳值，並且將回傳值傳入link參數
	for l := range c {
		counter++

		//當counter大於3時，關閉channel,這個不是一個好的方式,因為channel還有goroutine在執行
		if counter > 3 {
			fmt.Println("close channel")
			//關閉channel但是還有goroutine在執行
			//當goroutine執行完畢後傳送值到channel時，panic: send on closed channel
			close(c)
		}

		//建立goroutine，並且執行CheckLink()，並且傳入link參數
		//這裡的l是從channel取得的回傳值
		//因為channel是blocking的，所以這裡的goroutine會等待channel回傳值後，才會執行
		//這樣就可以達到非同步的效果
		go func(link string) {
			//每個link都會等待1秒後，再次執行CheckLink()
			time.Sleep(1 * time.Second)
			CheckLink(link, c)
		}(l)
	}

	time.Sleep(5 * time.Second)
}

func CheckLinkBlocking(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down! ****")
		return
	}
	fmt.Println(link, "is up! ****")
}

func CheckLink(link string, c chan string) {
	defer func() {
		c <- link
	}()

	rand := rand.Intn(10000)

	fmt.Println(link, "is checking...", rand)
	_, err := http.Get(link)
	time.Sleep(5 * time.Second)
	if err != nil {
		fmt.Println(link, "might be down!", rand)
		c <- link
		return
	}
	fmt.Println(link, "is up!", rand)
}
