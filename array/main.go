package main

import "fmt"

func main() {
	msg := [3]string{}
	msg[0] = "hello"
	msg[1] = "world"
	msg[2] = "golang"
	fmt.Println(msg)

	number := [...]int{1, 2, 3, 4, 5}
	fmt.Println(number)
	fmt.Println(number[1:4])
}
