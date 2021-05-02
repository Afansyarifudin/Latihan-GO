package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var angka = []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 25, 30, 45, 50}
	var min, max = getMinMax(angka)
	fmt.Print(" Data : ", angka, "\n")
	fmt.Print(" Nilai minimum : ", min, "\n")
	fmt.Print(" Nilai Maksimum : ", max, "\n")
}
