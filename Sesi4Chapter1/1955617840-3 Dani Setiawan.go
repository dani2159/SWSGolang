package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "selamat malam"
	kata := strings.Split(kalimat, " ")
	countKata := make(map[string]int)

	for _, k := range kata {
		for _, huruf := range k {
			fmt.Println(string(huruf))
			countKata[string(huruf)]++
		}
		fmt.Println()
	}
	fmt.Println(countKata)
}