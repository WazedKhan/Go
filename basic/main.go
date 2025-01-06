package main

import "fmt"

// with var key word we can declare variable outside the function
var CONSTANT string = "Hey I'm being called outside the function!"

// := declaration outside function is prohibited

func main() {
	name := "Snowy"
	age := 5
	fmt.Printf("Hello, %s!, You are now %d years old \n", name, age)
	fmt.Println(CONSTANT)
}
