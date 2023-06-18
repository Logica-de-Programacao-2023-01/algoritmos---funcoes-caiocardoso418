package main

import (
	"errors"
	"fmt"
)

func containsNumber(num int, nums []int) (bool, error) {
	if len(nums) == 0 {
		return false, errors.New("Erro: slice vazio")
	}

	for _, n := range nums {
		if n == num {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	num := 3

	result, err := containsNumber(num, nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("O número", num, "está presente no slice:", result)
	}

	emptySlice := []int{}

	result, err = containsNumber(num, emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("O número", num, "está presente no slice:", result)
	}
}
