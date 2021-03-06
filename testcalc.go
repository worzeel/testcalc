package testcalc

// Add numbers
func Add(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}

	return result
}

// Subtract numbers
func Subtract(numbers ...int) int {
	var result int

	for i, num := range numbers {
		if i == 0 {
			result = num
		} else {
			result -= num
		}
	}

	return result
}

// Multiply numbers
func Multiply(numbers ...int) int {
	var result int

	for i, num := range numbers {
		if i == 0 {
			result = num
		} else {
			result *= num
		}
	}

	return result
}
