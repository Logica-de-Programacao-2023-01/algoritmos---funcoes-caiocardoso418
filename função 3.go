package main

import "fmt"

func main() {
	juntar := []string{"fluminense", " ", "campeão"}
	fmt.Print(concatenacao(juntar))

}

func concatenacao(juntar []string) string {
	var resultado string

	for _, rs := range juntar {
		resultado += rs
	}
	return resultado

}
