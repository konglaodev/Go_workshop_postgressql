package main

import "fmt"

func main() {
	print("print")
	println("  println")
	fmt.Printf("fmt.Printf: value: %v, type: %T\n", 100, 100)
	dsn := fmt.Sprintf("host=%v port=%v", "localhost", 5432)
	print(dsn)
}
