package main

import "fmt"

func main() {
	// Deklarasi dan penugasan dalam satu baris
	namadepan := "nopal"
	namablkg := "kacong"
	var namatgh = "ucok"

	// string harus pakai +
	fmt.Println("Nama anda: " + namadepan + " " + namatgh + " " + namablkg)
	fmt.Printf("Nama anda: %s %s %s \n", namadepan, namatgh, namablkg)
}
