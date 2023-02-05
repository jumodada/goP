package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("filename.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
}
