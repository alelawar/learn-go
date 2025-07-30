package main

import "fmt"

func main() {
	// count := 1

	// for count <= 10000 {
	// 	fmt.Println("Pengulangan ke ", count)
	// 	count++
	// }

	// for count := 1; count <= 1000; count++ {
	// 	fmt.Println("Pengulangan ke ", count)
	// 	// count++
	// }

	names := []string{"ale", "kiki", "ali"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, name := range names {
		fmt.Println("Index: ", index, "name: ", name)
	}

}
