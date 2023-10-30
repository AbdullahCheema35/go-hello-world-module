package greeting

import (
	"time"
)

func GetCurrentGreeting() string {
	currentTime := time.Now()
	hour := currentTime.Hour()

	var greeting string

	switch {
	case hour >= 4 && hour < 12:
		greeting = "Good morning"
	case hour >= 12 && hour < 17:
		greeting = "Good afternoon"
	case hour >= 17 && hour < 20:
		greeting = "Good evening"
	case hour >= 20 && hour < 24:
		greeting = "Good night"
	case hour >= 0 && hour < 4:
		greeting = "Good late night"
	default:
		greeting = "Error 404"
	}

	return greeting
}
