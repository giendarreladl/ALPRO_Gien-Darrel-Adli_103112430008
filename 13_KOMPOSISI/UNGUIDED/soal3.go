package main

import (
	"fmt"
	"strings"
)

func main() {
	const percobaan = 5
	target := []string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 1; i <= percobaan; i++ {
		fmt.Printf("Percobaan %d: ", i)
		var gelas1, gelas2, gelas3, gelas4 string
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		input := []string{gelas1, gelas2, gelas3, gelas4}
		if strings.Join(input, " ") != strings.Join(target, " ") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
