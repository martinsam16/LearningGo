package main

import "fmt"

func main() {
	var meAma bool

	fmt.Println("Ella te ama?")
	fmt.Scan(&meAma)

	if meAma {
		fmt.Println("Estoy orgulloso de tí!")
	} else {
		fmt.Println("F")
	}
}
