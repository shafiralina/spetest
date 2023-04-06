package main

func BlueOceanReverse(arr1, arr2 []int) []int {
	arrlist := make(map[int]struct{}, len(arr2))
	for _, i := range arr2 {
		arrlist[i] = struct{}{}
	}
	var diff []int
	for _, i := range arr1 {
		if _, found := arrlist[i]; !found {
			diff = append(diff, i)
		}
	}
	return diff
}
