package main

import (
	"fmt"

	"github.com/AbdullahCheema35/go-hello-world-module/greeting"

	"github.com/AbdullahCheema35/go-hello-world-module/calculator"
)

func main() {
	greeting := greeting.GetCurrentGreeting()
	fmt.Printf("Hello World, %s!\n\n", greeting)

	fmt.Printf("1 + 2 = %d\n", calculator.Add(1, 2))
	fmt.Printf("2 - 1 = %d\n", calculator.Subtract(2, 1))
	fmt.Printf("2 * 3 = %d\n", calculator.Multiply(2, 3))
	fmt.Printf("6 / 2 = %d\n", calculator.Divide(6, 2))
}
