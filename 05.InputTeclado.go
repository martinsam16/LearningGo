package main

import "fmt"

func main() {
	var (
		nombre string
		edad   int8
	)

	fmt.Println("Ingresar nombres:")
	fmt.Scanln(&nombre)

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)

	fmt.Printf("Nombre: %s, Edad: %d", nombre, edad)

}
