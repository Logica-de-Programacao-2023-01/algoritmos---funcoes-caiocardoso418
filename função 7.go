package main

import (
	"errors"
	"fmt"
)

type IntFunction func(int) (int, error)

func applyFunctionToIntSlice(intSlice []int, fn IntFunction) ([]int, error) {
	if len(intSlice) == 0 {
		return nil, errors.New("Erro: slice vazio")
	}

	resultSlice := make([]int, len(intSlice))
	for i, num := range intSlice {
		result, err := fn(num)
		if err != nil {
			return nil, err
		}
		resultSlice[i] = result
	}

	return resultSlice, nil
}

func divisao(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Erro: divisão por zero")
	}
	return a / b, nil
}

func main() {
	intSlice := []int{10, 20, 30}
	divisionFn := divisao

	result, err := applyFunctionToIntSlice(intSlice, divisionFn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	emptySlice := []int{}
	result, err = applyFunctionToIntSlice(emptySlice, divisionFn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
