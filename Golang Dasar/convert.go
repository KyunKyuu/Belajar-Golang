package main

import "fmt"

func main() {
	var int32 int32 = 3231
	var int64 int64 = int64(int32)
	var int8 int8 = int8(int32)

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int8)

	stringg := "Teguh"
	t := stringg[0]
	convert := string(t)

	fmt.Println(stringg)
	fmt.Println(t)
	fmt.Println(convert)
}
