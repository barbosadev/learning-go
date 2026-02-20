package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//a := 10
	//b := 0
	//
	//res, err := divide(a, b)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res)
	//
	//user, _ := NewUser(true)
	//user.Foo()

	x := -10

	res, err := Sqrt(float64(x))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantError bool) (*User, error) {
	if wantError {
		return nil, errors.New("can't create user")
	}
	return &User{}, nil
}

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"can't square root of negative number"}
	}

	res := math.Sqrt(x)
	return res, nil
}
