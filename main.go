package main

import (
	"errors"
	"fmt"
)

type introduce func() map[string]string

type user struct {
	id   int
	name string
}

func (u *user) hi() map[string]string {
	return map[string]string{
		"name": u.name,
	}
}

var f = introduce(func() map[string]string {
	return map[string]string{
		"name": "Elon",
	}
})

func decorator(f introduce) introduce {
	return introduce(func() map[string]string {
		fmt.Println("Before call")
		result := f()
		fmt.Println("After call")
		return result
	})
}

func greet(f introduce) {
	fmt.Println("Hello", f()["name"])
}

func Max(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty slice")
	}

	var max int
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max, nil
}

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	fmt.Println("This is Design Pattern!")
}

type myType struct {
	name string
}

var global myType
var myInt int

func NewMyType() iface {
	m := myType{}
	global = m
	i := iface(&m)
	return i
}

func (m *myType) hi() string {
	return "hi " + m.name
}

func (m *myType) get() *myType {
	return m
}

type iface interface {
	hi() string
	get() *myType
}
