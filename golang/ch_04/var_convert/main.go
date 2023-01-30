package main

import (
	"fmt"
	"strconv"
)

func main() {
	//x := 3.8
	//var a int = int(x)
	//// 在go中， 不支持变量间的隐式类型转换
	//var c string = strconv.Itoa(3)
	//fmt.Printf("%T\n", c)
	//fmt.Printf("%T\n", a)

	//a2, err := strconv.Atoi("A")
	//if err != nil {
	//	fmt.Print(a2)
	//}
	//
	//a3, err2 := strconv.Atoi("C")
	//if err2 != nil {
	//	fmt.Print(a3)
	//}
	//fmt.Print(strconv.Atoi("1"))
	//b1, _ := strconv.ParseBool("true")
	//fmt.Printf("%b\n", b1)
	//
	//b2, _ := strconv.ParseFloat("3.1234567890", 1)
	//fmt.Println(b2)
	//
	//b3, _ := strconv.ParseInt("-40", 10, 64)
	//
	//fmt.Println(b3)

	b1 := strconv.FormatFloat(3.123456, 'f', -1, 64)
	b2 := strconv.FormatBool(true)

	fmt.Println(b1)

	fmt.Printf("%b\n", b1)
	fmt.Printf("%b\n", b2)
}
