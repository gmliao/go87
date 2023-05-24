package main

import (
	"fmt"
)

/*
	35 Testing and Benchmarking
	golang內建的測試框架,通常會在檔案名稱後面加上_test.go
	透過以下指令可以進行測試
	go test
	go test -v
	透過以下指令可以進行測試並且產生測試報告
	go test -coverprofile=coverage.out
	要測試的函式名稱必須以Test開頭,並且須以t *testing.T為參數
	譬如TestAdd(t *testing.T)
	測試函式中可以使用t.Error()來顯示錯誤訊息
	測試函式中可以使用t.Log()來顯示訊息
	測試函式中可以使用t.Fail()來顯示錯誤訊息並且結束測試
	要進行benchmark的函式名稱必須以Benchmark開頭,並且須以b *testing.B為參數
	譬如BenchmarkAdd(b *testing.B)
	透過以下指令可以進行benchmark
	go test -bench=.
    go test -bench=. -v
	透過以下指令可以進行benchmark並且產生benchmark報告
	go test -bench=. -benchmem
	透過以下指令可以進行benchmark並且產生benchmark報告,並且產生cpu profile
	go test -bench=. -benchmem -cpuprofile=cpu.out
	透過以下指令可以進行benchmark並且產生benchmark報告,並且產生memory profile
	go test -bench=. -benchmem -memprofile=mem.out
	透過以下指令可以進行benchmark並且產生benchmark報告,並且產生block profile
	go test -bench=. -benchmem -blockprofile=block.out
	透過以下指令可以進行benchmark並且產生benchmark報告,並且產生mutex profile
	go test -bench=. -benchmem -mutexprofile=mutex.out
	透過以下指令可以進行benchmark並且產生benchmark報告,並且產生trace profile
*/

func main() {
	user, err := FindUser(1)
	if err != nil {
		fmt.Printf("error finding user: %v\n", err)
		return
	}

	fmt.Printf("user: %v\n", user)

	err = user.UpdateName("Mike")
	if err != nil {
		fmt.Printf("error updating user: %v\n", err)
		return
	}
	fmt.Printf("user: %v\n", user)

	err = user.AddMoney(50)
	if err != nil {
		fmt.Printf("error adding money: %v\n", err)
		return
	}
	fmt.Printf("user: %v\n", user)
}
