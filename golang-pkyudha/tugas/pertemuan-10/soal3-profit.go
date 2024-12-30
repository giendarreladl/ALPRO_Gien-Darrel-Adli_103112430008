// ======= Pseudocode =======

/* kamus
    previous, current, difference : real
algoritma
    input(previous, current)
    difference = current - previous

    if difference > 0 then
        output("Peningkatan sebesar", difference)
    else if difference < 0 then
        output("Penurunan sebesar", -difference)
    else
        output("Tetap")
    endif
endprogram*/

// ======= Code =======
package main

import (
	"fmt"
)

func main() {
	var previous, current float64
	fmt.Println("Masukkan keuntungan bulan sebelumnya dan bulan ini (dipisahkan dengan spasi):")
	fmt.Scanf("%f %f", &previous, &current)

	difference := current - previous

	if difference > 0 {
		fmt.Printf("Peningkatan sebesar %.2f\n", difference)
	} else if difference < 0 {
		fmt.Printf("Penurunan sebesar %.2f\n", -difference)
	} else {
		fmt.Println("Tetap")
	}
}
