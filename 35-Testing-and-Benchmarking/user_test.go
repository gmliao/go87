package main

import (
	"testing"
)

// TestFindUser 測試 FindUser 函式
func TestFindUser(t *testing.T) {
	// 呼叫 FindUser 函數
	user, err := FindUser(1)

	// 確認函數是否有回傳錯誤
	if err != nil {
		t.Errorf("FindUser returned an error: %v", err)
	}

	// 確認 User 物件的 Id 是否正確
	if user.Id != 1 {
		t.Errorf("FindUser did not return the correct user")
	}
}

// TestUpdateName 測試 UpdateName 函式
func TestUpdateName(t *testing.T) {
	// 建立一個測試用的 User 物件
	user := &User{
		Name: "John",
		Age:  30,
	}

	// 呼叫 UpdateName 函式
	err := user.UpdateName("Mike")

	// 確認函式是否有回傳錯誤
	if err != nil {
		t.Errorf("UpdateName returned an error: %v", err)
	}

	// 確認 User 物件的 Name 是否已經被更新
	if user.Name != "Mike" {
		t.Errorf("UpdateName did not update the user's name")
	}
}

// TestAddMoney 測試 AddMoney 函式
func TestAddMoney(t *testing.T) {
	// 建立一個測試用的 User 物件
	user := &User{
		Name:  "John",
		Age:   30,
		Money: 100,
	}

	// 呼叫 AddMoney 函式
	err := user.AddMoney(50)

	// 確認函式是否有回傳錯誤
	if err != nil {
		t.Errorf("AddMoney returned an error: %v", err)
	}

	// 確認 User 物件的 Money 是否已經被更新
	if user.Money != 150 {
		t.Errorf("AddMoney did not update the user's money")
	}
}

// BenchmarkFindUser 測試 FindUser 函式的效能
func BenchmarkFindUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := FindUser(1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUpdateName 測試 UpdateName 函式的效能
func BenchmarkUpdateName(b *testing.B) {
	user := &User{Id: 1, Name: "User1", Age: 18}
	for i := 0; i < b.N; i++ {
		err := user.UpdateName("NewName")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkAddMoney 測試 AddMoney 函式的效能
func BenchmarkAddMoney(b *testing.B) {
	user := &User{Id: 1, Name: "User1", Age: 18, Money: 100}
	for i := 0; i < b.N; i++ {
		err := user.AddMoney(10)
		if err != nil {
			b.Fatal(err)
		}
	}
}
