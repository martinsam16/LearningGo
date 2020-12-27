package main

import "fmt"

func main() {
	mano()
	fmt.Println("Hola", dameComida())
	palabra, numero := dameMuchaComida()
	fmt.Println(palabra, numero)

	fmt.Println(sumar(3, 8))
}

func mano() {
	fmt.Println("Go pe mano 🌡")
}

// un solo retorno
func dameComida() string {
	return "Toma tu comida >:v"
}

// multiple retorno
func dameMuchaComida() (string, int) {
	return "Hola", 12
}

func sumar(a int, b int) int {
	return a + b
}
