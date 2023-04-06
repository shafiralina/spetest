package main

func NeedleHaystack(haystack []string, needle string) int {
	index := 0
	for i, data := range haystack {
		if data == needle {
			index = i
		}
	}

	return index
}
