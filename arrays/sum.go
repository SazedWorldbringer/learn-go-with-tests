package main

func Sum(numbrers [5]int) int {
	sum := 0

	for _, number := range numbrers {
		sum += number
	}

	return sum
}
