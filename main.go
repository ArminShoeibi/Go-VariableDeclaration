package main

import "fmt"

func main() {
	var firstName1 string = "Armin1"
	var lastName1 string = "Shoeibi1"

	var firstName2 = "Armin2"
	var lastName2 = "Shoeibi2"

	var (
		firstName3 = "Armin3"
		lastName3  = "Shoeibi3"
	)

	var firstName4, lastName4 = "Armin4", "Shoeibi4"

	firstName5 := "Armin5"
	lastName5 := "Shoeibi5"

	fmt.Println(firstName1, lastName1)
	fmt.Println(firstName2, lastName2)
	fmt.Println(firstName3, lastName3)
	fmt.Println(firstName4, lastName4)
	fmt.Println(firstName5, lastName5)
}
