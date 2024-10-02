package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	reamur := (4.0 / 5.0) * celsius
	fahrenheit := (celsius * 9.0 / 5.0) + 32
	kelvin := celsius + 273.15

	fmt.Println("Temperature Celcius:", celsius)
	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
