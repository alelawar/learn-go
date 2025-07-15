package main

import "fmt"

func main() {
	type NoKtp string

	var ktpAle NoKtp = "69696969"

	var contoh string = "46464646"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpAle)
	fmt.Println(contohKtp)
}
