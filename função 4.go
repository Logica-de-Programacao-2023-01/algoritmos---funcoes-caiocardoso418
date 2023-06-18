package main

import "fmt"

func main() {
	nums := []int{1, 5, 8, 2, 3}
	fmt.Print(segundomaiorvalor(nums))

}

func segundomaiorvalor(nums []int) int {
	var maior int
	var segundomaior int
	for _, rs := range nums {
		if rs > maior {
			segundomaior = maior
			maior = rs

		} else if rs > segundomaior && rs != maior {
			segundomaior = rs
		}
	}
	return segundomaior

}
