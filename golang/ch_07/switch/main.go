package main

import "fmt"

func main() {

	score := 72
	var grade string
	switch {
	case score > 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	case score >= 50:
		grade = "E"
	case score >= 40:
		grade = "F"
	default:
		grade = "W"
	}

	name := "Tom"
	switch name {
	case "Tim":
		grade = "A"
	case "Gai", "Tom":
		grade = "W"
	}

	fmt.Printf(grade)
}
