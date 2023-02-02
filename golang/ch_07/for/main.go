package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	name := "helloWorld"

	//for index, value := range name {
	//	fmt.Println(index)
	//	fmt.Printf("%c\n", value)
	//}

	name_arr := []rune(name)
	for i := 0; i < len(name_arr); i++ {
		fmt.Printf("%c\n", name[i])
	}
}
