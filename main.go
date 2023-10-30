package main

import (
	"fmt"

	"github.com/AbdullahCheema35/go-hello-world-module/greeting"
)

func main() {
	greeting := greeting.GetCurrentGreeting()
	fmt.Printf("Hello World, %s!\n", greeting)
}
