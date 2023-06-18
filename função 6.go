package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenateStrings(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("Erro: slice vazio")
	}

	result := strings.Join(stringsSlice, ",")
	return result, nil
}

func main() {
	stringsSlice := []string{"Olá", "Mundo", "Go"}

	result, err := concatenateStrings(stringsSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	emptySlice := []string{}

	result, err = concatenateStrings(emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
