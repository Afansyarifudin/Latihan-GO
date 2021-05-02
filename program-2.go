package main

import (
	"fmt"
	"math"
)

func lingkaran(d float64) (float64, float64) {
	//  menghitung luas lingkaran diketahui diameternya
	var luas = math.Pi * math.Pow(d/2, 2)
	// menghitung keliling
	var keliling = math.Pi * d

	// fungsi mengembalikan nilai luas dan keliling
	return luas, keliling

}

func main() {
	d := 15.00
	// memanggil fungsi lingkaran dengan 2 kembalian
	luas, keliling := lingkaran(d)
	fmt.Print(" \n")
	fmt.Print(" Diketahui diameter sebuah lingkaran adalah ", d, " \n")
	fmt.Print(" maka: \n")

	fmt.Print(" Luas lingkaran adalah ", luas, " \n")
	fmt.Print(" Keliling lingkaran adalah ", keliling, " \n")
}
