package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Age   int
	Money float64
}

// FindUser 透過id來查詢使用者
func FindUser(id int) (*User, error) {
	return &User{Id: id, Name: fmt.Sprintf("User%d", id), Age: 18}, nil
}

// UpdateName 透過name來更新使用者的名稱
func (u *User) UpdateName(name string) error {
	u.Name = name
	return nil
}

// AddMoney 透過money來增加使用者的金錢
func (u *User) AddMoney(money float64) error {
	u.Money += money
	return nil
}
