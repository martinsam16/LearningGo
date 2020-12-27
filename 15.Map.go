package main

import "fmt"

func main() {
	mapa := map[string]int{
		"key": 123,
		"xd":  25,
	}
	fmt.Println(mapa)

	mapa["other"] = 14
	fmt.Println(mapa)

	fmt.Println(mapa["key"])

	delete(mapa, "key")
	fmt.Println(mapa)

	mapa2 := make(map[int]string)
	fmt.Println(mapa2)
}
