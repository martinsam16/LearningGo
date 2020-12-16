package main

import "fmt"

func main() {
	// En go solo existe el for
	c := 0

	fmt.Println("Condicional")
	for c <= 5 {
		fmt.Println(c)
		c += 1
	}

	fmt.Println()

	fmt.Println("For normal")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		//break
		//continue
	}
}
