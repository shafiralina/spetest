package main

import "fmt"

func main() {
	number := 153
	fmt.Println(NarcissisticNumber(number))

	arr1 := []int{1, 5, 1, 3, 4}
	arr2 := []int{1}
	fmt.Println(BlueOceanReverse(arr1, arr2))

	arr3 := []int{2, 4, 6, 1}
	status, val := ParityOutlier(arr3)
	if status {
		fmt.Println(val)
	} else {
		fmt.Println(status)
	}

	needle := "blue"
	haystack := []string{"yellow", "green", "pink", "blue"}
	fmt.Println(NeedleHaystack(haystack, needle))
}
