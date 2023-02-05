package main

import "fmt"

func add(params ...int) (sum int) {
	for _, v := range params {
		sum += v
	}
	return sum
}

type sub func(a, b int) int

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(slice...))

	res := func(a int, b int) int {
		return a * b
	}(3, 2)

	fmt.Println(res)

}
