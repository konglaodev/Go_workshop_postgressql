package main

import (
	"./user"
	"fmt"
)

func main() {
	fmt.Printf("customer: name:%v, age:%v", user.Name, user.Age)
}
