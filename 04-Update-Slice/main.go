package main

import (
	"fmt"
)

/*
	slice的結構如下
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
	再傳遞的時候都是傳遞這個結構的指標
*/

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateArray(s [2]string) {
	s[0] = "Bye"
}

func main() {
	// golang裡面的 value type 有 int, float, bool, string, array
	// value type的特性是傳送的時候會複製一份，所以不會改變原本的值

	// golang裡面的 reference type 有 slice, map, channel, pointer, function
	// reference type的特性是傳送的時候會傳送一個指標，所以會改變原本的值

	// slice是reference type
	ss := []string{"Hello", "World"}
	updateSlice(ss)

	// 會印出[Bye World]
	fmt.Println(ss)

	// array是value type
	sa := [2]string{"Hello", "World"}
	updateArray(sa)

	// 會印出[Hello World]
	fmt.Println(sa)

	// slice的特性是可以動態增加長度
	foo := make([]int, 5)
	fmt.Println(foo, len(foo), cap(foo))

	// 增加長度
	foo = append(foo, 1)
	fmt.Println(foo, len(foo), cap(foo))

	// 當你使用 `foo[1:3]` 取出 `foo` 的部分值時，會建立一個新的 slice，這個新的 slice 是 `foo` 的一個 reference，
	// 所以當你改變 `bar` 的值時，`foo` 的值也會被改變。但是當你使用 `append` 增加 `foo` 的長度時，
	// 如果超過了 `foo` 的 `cap`，就會重新分配一塊更大的記憶體空間，這時 `foo` 的 reference 就會指向這個新的記憶體空間，
	// 而 `bar` 的 reference 仍然指向原本的記憶體空間，所以當你改變 `bar` 的值時，`foo` 的值不會被改變。
	bar := foo[1:3]
	bar[0] = 100
	fmt.Println(foo, len(foo), cap(foo))

	// 增加長度有可能會超過cap，所以會重新分配記憶體
	for i := 0; i < 100; i++ {
		foo = append(foo, 1)
	}

	bar[1] = 200
	fmt.Println(foo, len(foo), cap(foo))

	fmt.Println(bar, len(bar), cap(bar))

	// 如果要確保不會改變原本的值，可以使用copy
	copySlice := make([]int, 2)
	copy(copySlice, foo[1:3])

	s := []int{1, 2, 3}

	fmt.Println(s, len(s), cap(s))

	// 雖然有增加一筆資料,但是s還是指向舊的記憶體空間,不建議這種寫法
	appendToSlice1(s)
	fmt.Println(s, len(s), cap(s))

	// 把s的紀體空間的位址傳過去,所以會改變s的值,不建議這種寫法
	appendToSlice2(&s)
	fmt.Println(s, len(s), cap(s))

	// 建議這種寫法
	s = appendToSlice3(s)
	fmt.Println(s, len(s), cap(s))
}

func appendToSlice1(s []int) []int {
	s = append(s, 1)
	return s
}

func appendToSlice2(s *[]int) {
	*s = append(*s, 1)
}

func appendToSlice3(s []int) []int {
	return append(s, 1)
}
