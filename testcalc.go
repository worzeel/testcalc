package testcalc

func Add(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}

	return result
}
