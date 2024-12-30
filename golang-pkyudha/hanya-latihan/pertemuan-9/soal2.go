package main

import (
	"fmt"
	"strings"
)

func main() {
	var karakter string

	fmt.Scan(&karakter)

	karakter = strings.ToUpper(karakter)

	if karakter >= "A" && karakter <= "Z" {
		fmt.Print("Ini Alphabet")
	} else {
		fmt.Print("Bukan Alphabet")
	}
}
