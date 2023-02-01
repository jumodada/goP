package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "测试1"
	fmt.Println(len(name))

	var name1 = "hello world"
	strings.Contains(name1, "world")
	fmt.Println(strings.Index(name1, "world"))
	fmt.Println(strings.Count(name1, "l"))
	fmt.Println(strings.Compare("a", "a"))     //ASCII
	fmt.Println(strings.TrimSpace("   a"))     //ASCII
	fmt.Println(strings.TrimLeft(" c a", "c")) //ASCII
}
