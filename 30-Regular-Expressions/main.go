package main

import (
	"fmt"
	"regexp"
)

/*
	30 Regular Expressions
	這個章節主要是在講解正規表達式的使用，
	正規表達式是一種用來描述字串的方法，可以用來搜尋、取代、比對字串，
	這個範例介紹如何使用正規表達式，以及如何在 Go 語言中使用正規表達式。
*/

func main() {

	//生成一個身分整字號的正規表達式,開頭為[A-Z]後面接[12]後面接[0-9]{8}
	//這個正規表達式可以用來檢查身分證字號是否正確

	//直接使用matchString函式,缺點是每次都要重新編譯正規表達式
	match, _ := regexp.MatchString("^[A-Z][12][0-9]{8}$", "A123456789")

	fmt.Println(match)

	//使用MustCompile函式編譯正規表達式,這樣就可以重複使用
	//MustCompile函式會在編譯失敗時panic,所以不用檢查錯誤
	r := regexp.MustCompile("^[A-Z][12][0-9]{8}$")
	fmt.Println(r.MatchString("A123456789"))

	//性別錯誤
	fmt.Println(r.MatchString("A323456789"))
}
