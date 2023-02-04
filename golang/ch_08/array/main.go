package main

import "fmt"

func main() {
	test1 := [...]string{"11"}
	fmt.Println(test1)

	test2 := [...]int{1: 1, 3: 2, 10: 3}
	fmt.Println(test2)

	for value, index := range test1 {
		fmt.Println(value, index)
	}
}
