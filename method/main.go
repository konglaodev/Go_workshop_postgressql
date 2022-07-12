package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	name string
	age  int
}

func (c Customer) Hello() string {
	return "hello " + c.name
}

func (c *Customer) SetName() {
	c.name = "new"
}

func (c Customer) Validate() error {
	if c.name == "" {
		return errors.New("invalid")
	}
	return nil
}

func main() {
	c := Customer{name: "anousone", age: 23}
	fmt.Println(c)
	fmt.Println(c.name)
	c.SetName()
	fmt.Println(c.name)
	fmt.Println(c.Hello())

	if err := c.Validate(); err != nil {
		println("validate pass")
	}
}
