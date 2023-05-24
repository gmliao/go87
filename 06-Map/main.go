package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	// 第一種宣告方式
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
		"black": "#000000",
	}

	printMap(colors)

	// 第二種宣告方式
	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	colors2["black"] = "#000000"
	colors2["red"] = "#ff0000"

	printMap(colors2)

	// 刪除map中的元素
	delete(colors2, "white")
	printMap(colors2)
}
