package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	name := "helloWorld"

	for index, value := range name {
		fmt.Println(index)
		fmt.Printf("%c\n", value)
	}
}
