package main

import "fmt"

func main() {
	//slice1 := make([]string, 5)
	//slice2 := slice1[:]
	//slice2[0] = "44"
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//
	//subSlice := []int{1}
	//
	//subSlice2 := append(subSlice, 1)
	//fmt.Printf("%d\n", subSlice)
	//fmt.Printf("%d\n", subSlice2)
	//
	//subSlice3 := make([]int, 1)
	//copy(subSlice3, subSlice2)
	//fmt.Printf("%d\n", subSlice3)

	a := make([]int, 0)
	b := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(b)

	c := b[:]
	c[0] = 8
	fmt.Println(b)
	fmt.Println(c)
}
