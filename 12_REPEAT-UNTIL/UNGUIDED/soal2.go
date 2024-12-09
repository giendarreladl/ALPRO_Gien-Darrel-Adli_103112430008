package main

import "fmt"

func main() {
	var angka float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scanln(&angka)

	batas := float64(int(angka) + 1) // Membulatkan ke atas

	total := angka
	for {
		if total >= batas-0.1 {
			fmt.Printf("%d\n", int(batas))
			break
		}

		fmt.Printf("%.1f\n", total)
		total += 0.1
	}
}
