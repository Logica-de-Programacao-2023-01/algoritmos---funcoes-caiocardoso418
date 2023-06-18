package main

import (
	"errors"
	"fmt"
)

func divisao(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Erro: divisão por zero")
	}

	return a / b, nil
}

func main() {
	num1 := 10
	num2 := 2

	resultado, err := divisao(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	num2 = 0
	resultado, err = divisao(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
