package main

import "fmt"

func add(params ...int) (sum int) {
	for _, v := range params {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
}
