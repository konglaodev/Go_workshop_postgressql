package main

import "fmt"

func main() {
	score := 90
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	if err := IsError(); err != nil {
		println("error")
	}
}

func IsError() error {
	// return errors.New("some error")
	return nil
}
