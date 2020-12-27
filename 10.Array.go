package main

import "fmt"

func main() {
	var array [2]string
	array[0] = "hola"
	array[1] = "hola x2"
	fmt.Println(array)

	arrayNumbers := [3]int{1, 2, 3}
	fmt.Println(arrayNumbers)

	var soyArrayNumber [3]int
	soyArrayNumber[0] = 1
	soyArrayNumber[1] = 1
	soyArrayNumber[2] = 1
	fmt.Println(soyArrayNumber)
}
