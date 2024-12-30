//=======Pseudocode================================

/* kamus
    n, j, i : integer
algoritma
    input(n) // Masukkan bilangan ganjil

    for j = 1 to n do
        for i = 1 to n do
            if j = i or j = n - i + 1 then
                output(j)
            else
                output("  ") // Spasi kosong
            endif
        endfor
        output(baris baru)
    endfor
endprogram*/

//=======Code================================

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan ganjil: ")
	fmt.Scanln(&n)

	for j := 1; j <= n; j++ {
		for i := 1; i <= n; i++ {
			if j == i || j == n-i+1 {
				fmt.Printf("%d ", j)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
