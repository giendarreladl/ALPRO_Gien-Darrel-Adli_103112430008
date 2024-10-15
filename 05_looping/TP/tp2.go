package main

import (
	"fmt"
)

func main() {
	var n int

	// Meminta input
	fmt.Print("Masukkan jumlah baris segitiga bintang: ")
	fmt.Scan(&n)

	// Mencetak segitiga bintang
	for i := 1; i <= n; i++ {
		// Mencetak spasi untuk membuat segitiga di tengah
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// Mencetak bintang untuk membentuk segitiga simetris
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
