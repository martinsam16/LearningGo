package main

import "fmt"

func main() {
	var (
		x int
		y int
	)

	fmt.Println("Ingrese el valor de X")
	fmt.Scan(&x)

	fmt.Println("Ingrese el valor de Y")
	fmt.Scan(&y)

	fmt.Println("Calculando..")
	suma := x + y
	resta := x - y
	multiplicacion := x * y
	division := x / y
	mod := x % y

	fmt.Printf("X: %d Y: %d", x, y)
	fmt.Println()
	fmt.Println("SUMA			: ", suma)
	fmt.Println("RESTA			: ", resta)
	fmt.Println("MULTIPLICACIÓN	: ", multiplicacion)
	fmt.Println("DIVISIÓN		: ", division)
	fmt.Println("MOD			: ", mod)
}
