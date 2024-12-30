// ===========Pseudocode============

/* kamus
    a, b, c, d : integer
    n : integer
    max, min : integer
algoritma
    input(a, b, c, d)
    max = a
    min = a

    if n >= 2 then
        if b > max then
            max = b
        endif
        if b < min then
            min = b
        endif
    endif

    if n >= 3 then
        if c > max then
            max = c
        endif
        if c < min then
            min = c
        endif
    endif

    if n = 4 then
        if d > max then
            max = d
        endif
        if d < min then
            min = d
        endif
    endif

    output("Gol terbanyak:", max)
    output("Gol tersedikit:", min)
endprogram*/

// ===========Code============
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int

	fmt.Println("Masukkan jumlah gol (2 atau 4 tim, pisahkan dengan spasi):")
	n, _ := fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	max, min := a, a

	if n >= 2 {
		if b > max {
			max = b
		}
		if b < min {
			min = b
		}
	}
	if n >= 3 {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	if n == 4 {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
	}

	// Output hasil
	fmt.Printf("Gol terbanyak: %d\n", max)
	fmt.Printf("Gol tersedikit: %d\n", min)
}
