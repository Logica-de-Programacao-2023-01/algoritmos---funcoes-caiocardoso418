package main

import (
	"fmt"
)

func intersecao(x []int, y []int) ([]int, error) {

	intersecaoSlice := []int{}
	frequenciaNumX := make(map[int]int)
	frequenciaNumY := make(map[int]int)

	if len(x) == 0 || len(y) == 0 {
		return intersecaoSlice, fmt.Errorf("pelo menos um dos slices é vazio")
	}

	for _, ranX := range x {
		frequenciaNumX[ranX] += 1
	}

	for _, ranY := range y {
		frequenciaNumY[ranY] += 1
	}

	if len(x) >= len(y) {

		for i := 0; i < len(x); i++ {

			if frequenciaNumY[x[i]] != 0 {

				intersecaoSlice = append(intersecaoSlice, x[i])

			}

		}
	} else if len(x) < len(y) {

		for i := 0; i < len(y); i++ {

			if frequenciaNumX[y[i]] != 0 {

				intersecaoSlice = append(intersecaoSlice, y[i])

			}
		}
	}

	return intersecaoSlice, nil

}

func main() {

	x := []int{3423, 7, 8, 4, 5, 6}
	y := []int{3, 7, 5, 23425423432, 6, 2, 1}

	fmt.Printf("A interseção dos conjuntos %d com %d é:", x, y)
	fmt.Print(intersecao(x, y))

}
