package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	number = append(number, 6, 7, 8)
	fmt.Println(number)
	fmt.Println(number[:4])

	name := []string{}
	name = append(name, "anousone")
	name = append(name, "daky")
	fmt.Println(name)
}
