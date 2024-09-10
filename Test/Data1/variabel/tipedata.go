package main

import "fmt"

func main() {
	var angka int = 123

	var desimal float32 = 2.821

	// int harus pakai ,
	fmt.Printf("Angka: %d %.3f \n", angka, desimal)
	fmt.Println("Angka: %.5f", desimal)

	var benar bool = true
	var salah bool = false

	fmt.Println("Apakah: ", benar, " ", salah)
}
