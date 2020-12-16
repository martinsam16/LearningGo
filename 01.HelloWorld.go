package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Hola mundo desde GO 🥳 ")
	fmt.Println("¡Go p mano!")
	fmt.Println("Y que la fuerza te acompañe 🔦")

	fmt.Println()
	fmt.Println("⏰: ", time.Now())

	nombre := "Martín Alexis"
	edad := 20
	PI := 3.14

	fmt.Println()
	fmt.Printf("Nombre: %s, Edad: %d, PI: %f ", nombre, edad, PI)
}
