// ======= Pseudocode =======

/* kamus
    T, V, VolumeSekarang : integer
algoritma
    input(T)
    VolumeSekarang = 0

    while VolumeSekarang < T do
        input(V)
        VolumeSekarang = VolumeSekarang + V

        if VolumeSekarang >= T then
            output("true")
            break
        else
            output("false")
        endif
    endwhile
endprogram*/

//==========Code================================

package main

import "fmt"

func main() {
	var T int
	var V int
	fmt.Scan(&T)
	VolumeSekarang := 0

	for VolumeSekarang < T {

		fmt.Scan(&V)

		VolumeSekarang += V

		if VolumeSekarang >= T {
			fmt.Println("true")
			break
		} else {
			fmt.Println("false")
		}
	}
}
