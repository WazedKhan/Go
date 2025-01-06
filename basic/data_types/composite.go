package main

import "fmt"

// practice GO array
func array() {
	// Fixed-length array
	var arr = [5]int{1, 2, 3, 4, 5}

	// Array with inferred length
	var arrWithoutLen = [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println("Fixed-length array:", arr)
	fmt.Println("Inferred-length array:", arrWithoutLen)

	assign_array := [5]int{1: 10, 2: 40}
	fmt.Println("initializes only the second and third elements of the array", assign_array)
}

func main() {
	array()
}
