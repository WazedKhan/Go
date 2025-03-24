package main

import (
	"fmt"

	"practice/data_types" // Importing local package
)

func main() {
	fmt.Println("===== Structs Example =====")

	car := data_types.Car{
		Brand: "Toyota", Model: "Corolla", Year: 2019, FontWheel: data_types.Wheel{Radius: 15, Material: "Alloy"}, BackWheel: data_types.Wheel{Radius: 15, Material: "Alloy"}}
	fmt.Println(data_types.CarInfo(car))
}
