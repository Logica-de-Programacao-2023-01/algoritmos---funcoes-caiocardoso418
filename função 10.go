package main

import (
	"errors"
	"fmt"
	"sort"
)

func sortIntSlice(intSlice []int) ([]int, error) {
	if len(intSlice) == 0 {
		return nil, errors.New("Erro: slice vazio")
	}

	sortedSlice := make([]int, len(intSlice))
	copy(sortedSlice, intSlice)

	sort.Ints(sortedSlice)

	return sortedSlice, nil
}

func main() {
	intSlice := []int{9, 2, 5, 1, 7}

	result, err := sortIntSlice(intSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	emptySlice := []int{}
	result, err = sortIntSlice(emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
