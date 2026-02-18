package main

import (
	"encoding/json"
	"fmt"
)

type MyString string

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u User) PrintName() {
	fmt.Println("I am", u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	user := User{"Victor", 10}
	fmt.Println(user)
	user.UpdateName("Bob")
	user.PrintName()

	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
