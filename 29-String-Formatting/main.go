package main

import (
	"fmt"
	"os"
)

/*
	https://gobyexample.com/string-formatting
	這個範例主要是展示如何使用 fmt.Printf 來格式化字串
*/

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// 這邊的 %v 會將 struct 的內容印出來
	// {1 2}
	fmt.Printf("struct1: %v\n", p)

	// %+v 會將 struct 的內容印出來，並且會將欄位名稱也印出來
	// {x:1 y:2}
	fmt.Printf("struct2: %+v\n", p)

	// %#v 會將 struct 的內容印出來，並且會將欄位名稱也印出來，並且會將 struct 的型別也印出來
	// main.point{x:1, y:2}
	fmt.Printf("struct3: %#v\n", p)

	// %T 會將 struct 的型別印出來
	// main.point
	fmt.Printf("type: %T\n", p)

	// %t 會將 bool 值印出來, true 會印出 true, false 會印出 false
	// true
	fmt.Printf("bool: %t\n", true)

	// %d 會將 int 值印出來
	// 123
	fmt.Printf("int: %d\n", 123)

	// %b 會將 int 值印出來，並且會將 int 值轉成二進位
	// 1110
	fmt.Printf("bin: %b\n", 14)

	// %c 會將 int 值印出來，並且會將 int 值轉成 ASCII 字元
	// char: !
	fmt.Printf("char: %c\n", 33)

	// %x 會將 int 值印出來，並且會將 int 值轉成十六進位
	// hex: 1c8
	fmt.Printf("hex: %x\n", 456)

	// %f 會將 float 值印出來
	// float1: 78.900000
	fmt.Printf("float1: %f\n", 78.9)

	// %e 會將 float 值印出來，並且會將 float 值轉成科學記號
	// float2: 1.234000e+08
	fmt.Printf("float2: %e\n", 123400000.0)

	// %E 會將 float 值印出來，並且會將 float 值轉成科學記號，並且會將 e 轉成 E
	// float3: 1.234000E+08
	fmt.Printf("float3: %E\n", 123400000.0)

	// %s 會將 string 值印出來
	// str1: "string"
	fmt.Printf("str1: %s\n", "\"string\"")

	// %q 會將 string 值印出來，並且會將 string 值加上雙引號
	// str2: "\"string\""
	fmt.Printf("str2: %q\n", "\"string\"")

	// %x 會將 string 值印出來，並且會將 string 值轉成十六進位
	// str3: 6865782074686973
	fmt.Printf("str3: %x\n", "hex this")

	// %p 會將 pointer 值印出來
	// pointer: 0xc00000a0b0 (每次執行的結果都不一樣)
	fmt.Printf("pointer: %p\n", &p)

	// %6d 會將 int 值印出來，並且會將 int 值的寬度設為 6
	// width1: |    12|   345|
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// %6.2f 會將 float 值印出來，並且會將 float 值的寬度設為 6，並且會將 float 值的小數點後的位數設為 2
	// width2: |  1.20|  3.45|
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// %-6.2f 會將 float 值印出來，並且會將 float 值的寬度設為 6，並且會將 float 值的小數點後的位數設為 2，並且會將 float 值靠左對齊
	// width3: |1.20  |3.45  |
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// %6s 會將 string 值印出來，並且會將 string 值的寬度設為 6
	// width4: |   foo|     b|
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// %-6s 會將 string 值印出來，並且會將 string 值的寬度設為 6，並且會將 string 值靠左對齊
	// width5: |foo   |b     |
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// sprintf 會將結果輸出s變數中
	// %q 會將 string 值印出來，並且會將 string 值加上雙引號，並且會將 string 值的寬度設為 6
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Fprintf 會將結果輸出到 w 變數中, 必須要傳入一個 io.Writer 介面
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
