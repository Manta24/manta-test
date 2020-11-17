package main

import "fmt"

func main() {
	fmt.Println("C'est qui le plus beau")

	chou()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {

			fmt.Println(i)
		}
	}
}

func chou() {
	fmt.Println("C'est chouchou")
}

// Control Flow :
// (1) Sequence
// (2) loop; iterative
// (3) conditional
