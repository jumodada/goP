package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
func main() {

	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

	arr1 := [3]int{1, 2, 3}
	var arr2 *[3]int = &arr1
	fmt.Printf("%d\n", *arr2)
}
