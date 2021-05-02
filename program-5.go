package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"Nur", "Afan", "Syarifudin"}
	var dataContainsu = filter(data, func(each string) bool {
		return strings.Contains(each, "u")
	})

	var dataContainsA = filter(data, func(each string) bool {
		return strings.Contains(each, "A")
	})

	var dataLength4 = filter(data, func(each string) bool {
		return len(each) == 4
	})

	fmt.Print(" \n")
	fmt.Print(" Data awal : ", data, "\n")
	fmt.Print(" Data yang terdapat huruf u : ", dataContainsu, "\n")
	fmt.Print(" Data yang terdapat huruf Kapital A : ", dataContainsA, "\n")
	fmt.Print(" Data dengan jumlah huruf 4 : ", dataLength4, "\n")

}
