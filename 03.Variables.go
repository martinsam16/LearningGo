package main

import "fmt"

func main() {

	//
	var nombre string
	nombre = "Martín Alexis"

	fmt.Println(nombre)

	//
	var uno, dos, tres int

	fmt.Println(uno, dos, tres)

	//
	var (
		pi     float64
		meAma  bool
		cadena = "Chain!"
		edad   = 20
	)
	pi = 3.1492
	meAma = true

	fmt.Println(pi, meAma, cadena, edad)

	//
	sinVarNiTipoDeDato := 40
	sinVarNiTipoDeDato = 20
	fmt.Println(sinVarNiTipoDeDato)
}
