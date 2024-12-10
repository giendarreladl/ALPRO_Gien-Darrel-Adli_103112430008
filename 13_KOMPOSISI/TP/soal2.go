package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	faktor := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			faktor = append(faktor, i)
			jumlah += i
		}
	}

	fmt.Printf("Hasil: ")
	if jumlah == n {
		fmt.Printf("Ya (karena faktor dari %d adalah %v dan %d = %d)\n", n, faktor, jumlah, n)
	} else {
		fmt.Printf("Tidak (karena faktor dari %d adalah %v dan %d ≠ %d)\n", n, faktor, jumlah, n)
	}
}
