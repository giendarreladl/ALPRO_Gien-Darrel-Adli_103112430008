package main

import (
	"fmt"
	"strings"
)

func main() {
	var pita, bunga string
	count := 0

	for {
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scan(&bunga)

		if strings.EqualFold(bunga, "SELESAI") {
			break
		}

		if count > 0 {
			pita += " - "
		}
		pita += bunga
		count++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count)
}
