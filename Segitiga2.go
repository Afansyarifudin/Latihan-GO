package main

import "fmt"

func main() {
	var (
		n int = 7
	)

	fmt.Printf(" \n Program Membuat Pola Segitiga 2 dengan inputan n \n\n")
	fmt.Printf("\n")
	for i := n; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k <= n-i; k++ {
			fmt.Printf("*")
		}
		for k := 1; k < n-i; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
