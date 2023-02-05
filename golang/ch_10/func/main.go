package main

import "fmt"

func add(params ...int) (sum int) {
	for _, v := range params {
		sum += v
	}
	return sum
}

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(slice...))
}
