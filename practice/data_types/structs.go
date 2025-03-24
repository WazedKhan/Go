// data_types/structs.go
package data_types

import "fmt"

// Exported struct
type Car struct {
	Brand string
	Model string
	Year  int
	FontWheel Wheel
	BackWheel Wheel
}

type Wheel struct {
	Radius int
	Material  string
}

// Exported function
func CarInfo(car Car) string {
	return fmt.Sprintf("Car Information:\nBrand: %s\nModel: %s\nYear: %d\nFront Wheel:\n  Radius: %d\n  Material: %s\nBack Wheel:\n  Radius: %d\n  Material: %s", 
		car.Brand, car.Model, car.Year, car.FontWheel.Radius, car.FontWheel.Material, car.BackWheel.Radius, car.BackWheel.Material)
}
