package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hoooolaa amigoos", "o"))
	fmt.Println(strings.Contains("Hoooolaa amigoos", "ñ"))
	fmt.Println(strings.Count("re gil el wacho", "e"))
	fmt.Println(strings.ToUpper("hola xd"))
	fmt.Println(strings.ToLower("HOLA XD"))
	fmt.Println(strings.ReplaceAll("hola amix como estan", "a", "*"))
}
