package main

func oddOrEven(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func createIntSlice(number int) []int {
	start := 0
	newSlice := []int{}
	for start < number {
		newSlice = append(newSlice, start)
		start++
	}
	return newSlice
}
