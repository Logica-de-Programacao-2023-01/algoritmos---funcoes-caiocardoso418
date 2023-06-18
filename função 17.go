package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func sortStrings(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("Erro: slice vazio")
	}

	sort.Strings(stringsSlice)
	result := strings.Join(stringsSlice, "")

	return result, nil
}

func main() {
	stringsSlice := []string{"banana", "laranja", "abacaxi"}

	result, err := sortStrings(stringsSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", result)
	}

	emptySlice := []string{}

	result, err = sortStrings(emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
