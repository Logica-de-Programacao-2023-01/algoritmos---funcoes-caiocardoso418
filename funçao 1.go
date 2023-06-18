package main

import (
	"fmt"
)

func calcularMedia(valores []int) float64 {
	soma := 0
	for _, valor := range valores {
		soma += valor
	}
	media := float64(soma) / float64(len(valores))
	return media
}

func main() {
	valores := []int{6, 9, 64, 83, 101}

	media := calcularMedia(valores)
	fmt.Printf("A média é: %.2f\n", media)
}
