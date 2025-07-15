package main

import "fmt"

func main() {
	// const name1 = "ale"
	// const name2 = "ali"

	// const res bool = len(name1) == len(name2)

	// fmt.Println(res)

	const nilaiTotal = 80
	const absensi = 90

	const minimalNilai bool = nilaiTotal > 80
	const minimalAbsen bool = absensi > 80

	const lulus bool = minimalAbsen && minimalNilai

	fmt.Println(lulus)
}
