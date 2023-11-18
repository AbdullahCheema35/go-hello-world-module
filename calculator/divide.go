package calculator

func Divide(x int, y int) int {
	// Divide by repeated subtraction
	quotient := 0
	for x >= y {
		x = Subtract(x, y)
		quotient++
	}
	return quotient
}
