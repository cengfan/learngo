package main

import "time"

func main() {
	i := 2
	switch i {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")

	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("<12")
	default:
		println(">=12")
	}
}
