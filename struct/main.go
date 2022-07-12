package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func main() {
	c := Customer{name: "anousone", age: 23}
	fmt.Println(c)
	fmt.Println(c.name)
	fmt.Println(c.age)
}
