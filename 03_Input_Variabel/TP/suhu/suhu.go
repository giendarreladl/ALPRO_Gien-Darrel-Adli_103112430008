package main

import (
	"fmt"
)

func main() {

	var f float32
	var rumus float32

	fmt.Println("Masukan suhu farenheit :")
	fmt.Scanln(&f)
	rumus = (f * 5 / 9) + 273
	fmt.Println("Suhu kelvin nya adalah :", rumus)
}
