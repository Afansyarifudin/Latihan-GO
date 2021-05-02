package main

import (
	"fmt"
)

func main() {
	var (
		n      int
		i      int
		nama   [10]string
		status [10]string
		nilai  [10]string
	)

	fmt.Print("Masukan jumlah data : ")
	fmt.Scan(&n)
	fmt.Print(" \n")

	for i := 0; i < n; i++ {
		fmt.Print(" \n")

		fmt.Print(" Data ke-", i+1, " \n")
		fmt.Print(" Masukan nama : ")
		fmt.Scan(&nama[i])
		fmt.Print(" Masukan nilai : ")
		fmt.Scan(&nilai[i])

		if nilai[i] <= "78" && nilai[i] >= "60" {
			status[i] = "Remidial"
		} else if nilai[i] <= "59" {
			status[i] = "Tidak Lulus"
		} else {
			status[i] = "Lulus"
		}
	}

	fmt.Print(" \n")
	fmt.Print(" Daftar Nilai Praktikum Mahasiswa \n")
	fmt.Print(" ------------------------------------------- ", " \n")
	fmt.Print(" No    Nama             Nilai         Status          ", " \n")
	fmt.Print(" ------------------------------------------- ", " \n")

	for i = 0; i < n; i++ {
		fmt.Print(" ", i+1, "   ", nama[i], "             ", nilai[i], "            ", status[i], " \n")
		fmt.Print(" ------------------------------------------- ", " \n")
	}
}
