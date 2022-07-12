package main

import "fmt"

func main() {
	greeting := map[int]string{1: "hi", 2: "hello"}
	greeting[3] = "sabaidee"
	fmt.Println(greeting)
	fmt.Println(greeting[1])
}
