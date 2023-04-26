package main

import "fmt"

func main() {
	var nama string
	nama = "pake type string"
	fmt.Println(nama)

	var namaA = "tanpa type string"
	fmt.Println(namaA)

	angka := 23
	fmt.Println(angka)

	var (
		namaB = "Test variable banyak sekaligus"
		namaC = "Test variable banyak sekaligus"
	)

	fmt.Println(namaB)
	fmt.Println(namaC)

	var (
		angka1 int
		angka2 uint
	)

	angka1 = 2
	angka2 = 10

	fmt.Println(angka1)
	fmt.Println(angka2)

}
