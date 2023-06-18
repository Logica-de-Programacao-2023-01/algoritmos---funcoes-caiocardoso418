package main

import (
	"errors"
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func getPrimes(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("Erro: número negativo")
	}

	primes := []int{}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

func main() {
	num := 20

	primes, err := getPrimes(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Números primos:", primes)
	}

	negativeNum := -5

	primes, err = getPrimes(negativeNum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Números primos:", primes)
	}
}
