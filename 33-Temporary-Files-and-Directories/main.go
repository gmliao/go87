package main

import (
	"fmt"
	"os"
)

/*
	33 Temporary Files and Directories
		Golang 內建的 os 包提供了一些方法來處理臨時檔案和目錄
		1. func TempFile(dir, prefix string) (f *os.File, err error)
			- 在 dir 目錄下創建一個新的臨時檔案，名稱以 prefix 為前綴，返回臨時檔案的指針
		2. func TempDir(dir, prefix string) (name string, err error)
			- 在 dir 目錄下創建一個新的臨時目錄，名稱以 prefix 為前綴，返回臨時目錄的名稱

*/

func main() {

	// 創建臨時檔案
	tmpFile, err := os.CreateTemp("", "temp")
	if err != nil {
		panic(err)
	}

	fmt.Println(tmpFile.Name())

	tmpFile.Write([]byte("This is a temporary file"))

	tmpFile.Close()

	// 結束後刪除臨時檔案
	defer os.Remove(tmpFile.Name())

	// 創建臨時目錄
	tmpDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		panic(err)
	}

	fmt.Println(tmpDir)

	// 結束後刪除臨時目錄
	defer os.RemoveAll(tmpDir)
}
