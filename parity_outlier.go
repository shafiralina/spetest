package main

func ParityOutlier(arr []int) (bool, int) {
	var data_temp []int
	even := 0
	for _, data := range arr {
		if data%2 != 0 {
			data_temp = append(data_temp, data)
		} else {
			even = data
		}
	}
	if len(data_temp) == 1 {
		return true, data_temp[0]
	} else if len(data_temp) == len(arr)-1 {
		return true, even
	} else {
		return false, 0
	}
}
