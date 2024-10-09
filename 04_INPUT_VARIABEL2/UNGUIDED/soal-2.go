package main

import "fmt"

func main() {
	var bmi, tinggibadan float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scanln(&bmi)

	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scanln(&tinggibadan)

	beratBadan := bmi * tinggibadan * tinggibadan

	fmt.Printf("Berat badan Anda adalah: %.0f kg\n", beratBadan)
}
