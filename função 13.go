package main

import (
	"errors"
	"fmt"
)

func sumDigits(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Erro: número negativo")
	}

	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}

	return sum, nil
}

func main() {
	num := 12345

	result, err := sumDigits(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Soma dos dígitos:", result)
	}

	negativeNum := -123

	result, err = sumDigits(negativeNum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Soma dos dígitos:", result)
	}
}
