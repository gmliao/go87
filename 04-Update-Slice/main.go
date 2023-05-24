package main

import fm "fmt"

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
	fm.Println(ss)

	// array是value type
	sa := [2]string{"Hello", "World"}
	updateArray(sa)

	// 會印出[Hello World]
	fm.Println(sa)
}
