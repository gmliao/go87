package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	31 Line Filters
	利用scanner讀取標準輸入的每一行,並將每一行的字串轉換成大寫後輸出
	windows利用下面指令可以把檔案輸出到標準輸入,並且傳送到程式
		type hello.txt | go run .
	linux
		cat hello.txt | go run .
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if scanner.Err() != nil {
		fmt.Println("error:", scanner.Err())
		os.Exit(1)
	}
}
