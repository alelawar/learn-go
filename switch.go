package main

import "fmt"

func main() {
	name := "jamil"

	switch name {
	case "ale":
		fmt.Println("iyah ale")
	case "kiki":
		fmt.Println("iyah kiki")
	default:
		fmt.Println("Kenalan dlu der")
	}

	// length := len(name) > 5

	switch length := len(name); length > 1 {
	case true:
		fmt.Println("Kepanjangan")
	case false:
		fmt.Println("Kependekan")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("kepanjangan")
	case length > 5:
		fmt.Println("Mayan")
	default:
		fmt.Println("Udah oke")
	}
}
