package main

import "fmt"

func main() {
	var numero int8

	fmt.Println("Ingrese un número:")
	fmt.Scan(&numero)

	switch numero {
	case 1:
		fmt.Println("UNO")
		break
	case 2:
		fmt.Println("DOS")
		break
	default:
		fmt.Println("No hay excusa 🐙")
		break
	}
}
