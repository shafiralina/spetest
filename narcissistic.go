package main

import (
	"math"
)

func NarcissisticNumber(num int) bool {
	var numslice []int
	num_temp := num
	for num_temp > 0 {
		numslice = append(numslice, num_temp%10)
		num_temp = num_temp / 10
	}
	amount := 0
	for _, i := range numslice {
		amount = amount + int(math.Pow(float64(i), float64(len(numslice))))
	}

	return amount == num
}
