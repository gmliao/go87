package main

import (
	"fmt"
	"net/http"
)

//channel使用時機,goroutine之間需要溝通

func main() {
	cGoogle := make(chan string)
	cFacebook := make(chan string)
	cStackoverflow := make(chan string)

	go CheckLink("http://google.com", cGoogle)
	go CheckLink("http://facebook.com", cFacebook)
	go CheckLink("http://stackoverflow.com", cStackoverflow)

	//有三個channel,所以要做三次select
	//select會等待channel回傳值,只要有一個channel回傳值,就會執行case
	for i := 0; i < 3; i++ {
		select {
		case googleStatus := <-cGoogle:
			fmt.Println(googleStatus)
		case facebookStatus := <-cFacebook:
			fmt.Println(facebookStatus)
		case stackoverflowStatus := <-cStackoverflow:
			fmt.Println(stackoverflowStatus)
		}
	}

}
func CheckLink(link string, c chan string) {
	//golang特性,defer會在函式結束時執行,所以可以用來做一些收尾工作
	//要把channel傳進去defer,不然defer會有問題
	defer func(c chan string) {
		c <- link
	}(c)
	_, err := http.Get(link)
	if err != nil {
		c <- link
		return
	}
	fmt.Println(link, "is up!")
}
