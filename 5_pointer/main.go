package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	// 0xc000052270
	fmt.Println(&name)
	// 0xc000052270
	fmt.Println(namePointer)
	// 0xc00000a028
	fmt.Println(&namePointer)

	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	// 0xc00000a038
	fmt.Println(&namePointer)
	// 0xc000052270
	fmt.Println(namePointer)
}
