package main

import "fmt"

func main() {
	var beratGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratGram)

	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := beratKg * 10000
	biayaSisa := 0

	if beratKg <= 10 {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", biayaKg+biayaSisa)
}
