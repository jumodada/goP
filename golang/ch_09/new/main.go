package main

import "fmt"

func main() {
	c := new(int)
	fmt.Println(*c)

	d := make(map[string]string)

	d["1"] = "1"
}
