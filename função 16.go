package main

import (
	"errors"
	"fmt"
	"strings"
)

func replaceVowels(str string) (string, error) {
	if str == "" {
		return "", errors.New("Erro: string vazia")
	}

	vowels := "aeiouAEIOU"
	replacement := "*"

	result := strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return []rune(replacement)[0]
		}
		return r
	}, str)

	return result, nil
}

func main() {
	input := "Hello, World!"
	result, err := replaceVowels(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", result)
	}

	emptyString := ""
	result, err = replaceVowels(emptyString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
