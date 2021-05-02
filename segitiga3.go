package main

import "fmt"

func main() {
	var (
		n int = 7
	)

	fmt.Printf(" Program Membuat Pola Segita 3 dengan Inputan n \n\n")
	fmt.Printf("\n")
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			fmt.Printf("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf(" *")
		}
		fmt.Printf("\n")
	}
}
