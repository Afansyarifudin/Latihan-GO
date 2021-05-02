package main

import "fmt"

func segitiga(x, y int) int {
	L := (x * y) / 2
	return L
}
func main() {
	var (
		a, t int
	)
	fmt.Print(" Masukan alas segitiga : ")
	fmt.Scan(&a)
	fmt.Print(" Masukan tinggi segitiga : ")
	fmt.Scan(&t)
	fmt.Print(" ------------------------- \n")
	fmt.Print(" Luas segitiga adalah ", segitiga(a, t), "\n")

}
