package main

import "fmt"

func main() {
	// names := [...]string{
	// 	"ale",
	// 	"kiki",
	// 	"ali",
	// 	"padlan",
	// }

	// slice1 := names[0:3]
	// fmt.Println(slice1)

	// slice2 := names[2:4]
	// fmt.Println(slice2)

	// var slice3 []string = names[3:]
	// fmt.Println(slice3)
	// fmt.Println("Length of slice3:", len(slice3))

	// slice4 := names[:]
	// fmt.Println(slice4)
	// fmt.Println("Cap of slice4: ", cap(slice2))

	// days := [...]string{
	// 	"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	// }

	// slice1 := days[5:]        // sabtu, minggu
	// slice1[0] = "sabtu baru"  // kita tambahin data baru
	// slice1[1] = "minggu baru" // kita tambahin data baru
	// fmt.Println(slice1)
	// fmt.Println(days) // data baru dari slice yg diatas

	// daySlice2 := append(slice1, "mingse")
	// daySlice2[0] = "sabtu lama"
	// // daysBaru := [...]string{
	// //	"senin", "selasa", "rabu", "kamis", "jumat", "sabtu baru", "minggu baru", "mingse",
	// // } Dia buat array baru dari si slice1
	// fmt.Println(slice1)
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "ale"
	newSlice[1] = "kiki"
	// newSlice[2] = "padlan" Error karna maksimal length yg bisa di timpa cuma sampe 2

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	newSlice2 := append(newSlice, "padlan")
	fmt.Println(newSlice2)
	fmt.Println(cap(newSlice2))
	fmt.Println(len(newSlice2))
	newSlice2[0] = "alelawar"

	fmt.Println(newSlice2)
	fmt.Println(newSlice)
}
