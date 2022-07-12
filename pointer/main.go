package main

import "fmt"

func main() {
	x := 10
	var y *int
	y = &x
	*y = 20
	fmt.Printf("x:%v\n", x)
	fmt.Printf("y:%v\n", *y)
	myAge := 23
	update(&myAge)
	fmt.Println(myAge)
}

func update(age *int) {
	*age = *age - 5
}
