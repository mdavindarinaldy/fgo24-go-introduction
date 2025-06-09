package main

import (
	"fgo24-golang-basic/utils"
	"fmt"
)

// func area(r float64) float64 {
// 	phi := 3.14 // karena tidak dideklarasikan secara explisit, dia tipe datanya float yg default yaitu float64
// 	return phi * r * r
// }

// func circumference(r float64) float64 {
// 	phi := 3.14
// 	return 2 * phi * r
// }

func main() {
	// fmt.Println(area(7))
	// fmt.Println(circumference(7))
	fmt.Println(utils.Area(7))
	fmt.Println(utils.Circumference(7))
}
