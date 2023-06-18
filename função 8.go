package main

import (
	"errors"
	"fmt"
)

func filterEvenNumbers(intSlice []int) ([]int, error) {
	if len(intSlice) == 0 {
		return nil, errors.New("Erro: slice vazio")
	}

	evenNumbers := make([]int, 0)
	for _, num := range intSlice {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}

	return evenNumbers, nil
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result, err := filterEvenNumbers(intSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	emptySlice := []int{}
	result, err = filterEvenNumbers(emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
