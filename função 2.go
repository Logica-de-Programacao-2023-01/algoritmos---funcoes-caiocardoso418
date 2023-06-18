package main

import "fmt"

func main() {
	var x string
	fmt.Println("diga uma palavra:")
	fmt.Scan(&x)
	fmt.Println(quantidadevogais(x))

}

func quantidadevogais(x string) int {
	contagem := 0

	for _, rs := range x {
		switch rs {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			contagem++
		}
	}
	return contagem
}
