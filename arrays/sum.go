package main

func Sum(numbrers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbrers[i]
	}
	return sum
}
