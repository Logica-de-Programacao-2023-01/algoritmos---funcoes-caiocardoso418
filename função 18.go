package main

import (
	"errors"
	"fmt"
)

func applyFunctionAndSum(numbers []int, f func(int) int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("Erro: slice vazio")
	}

	sum := 0
	for _, num := range numbers {
		result := f(num)
		sum += result
	}

	return sum, nil
}

func multiplyByTwo(num int) int {
	return num * 2
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	result, err := applyFunctionAndSum(numbers, multiplyByTwo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Soma:", result)
	}

	emptySlice := []int{}

	result, err = applyFunctionAndSum(emptySlice, multiplyByTwo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Soma:", result)
	}
}
