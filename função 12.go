package main

import (
	"errors"
	"fmt"
	"strings"
)

func filterUpperCase(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("Erro: slice vazio")
	}

	var result strings.Builder
	for _, str := range stringsSlice {
		if len(str) > 0 && str[0] >= 'A' && str[0] <= 'Z' {
			result.WriteString(str)
			result.WriteString(",")
		}
	}

	return strings.TrimSuffix(result.String(), ","), nil
}

func main() {
	stringsSlice := []string{"Olá", "Mundo", "Go", "Programação"}

	result, err := filterUpperCase(stringsSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	emptySlice := []string{}

	result, err = filterUpperCase(emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
