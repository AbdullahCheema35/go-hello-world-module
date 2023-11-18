package calculator

func Multiply(x int, y int) int {
	// Multiply by adding x to itself y times
	product := 0
	for i := 0; i < y; i++ {
		product = Add(product, x)
	}
	return product
}
