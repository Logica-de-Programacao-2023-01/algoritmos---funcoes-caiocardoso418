package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrimeNumber(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("Erro: número negativo")
	}

	if num < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(num)))
	for i := 2; i <= limite; i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	num := 17

	result, err := isPrimeNumber(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	negativeNum := -5
	result, err = isPrimeNumber(negativeNum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
