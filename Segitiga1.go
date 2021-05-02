package main

import "fmt"

func main() {
	var (
		n int = 7
	)

	fmt.Printf("\n Program Membuat Pola Segitiga 1 dengan Inputan n  \n\n")
	fmt.Printf("\n")
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" *")
		}
		fmt.Printf("\n")
	}
}
