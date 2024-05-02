package main

import "fmt"

func main() {
	// default to 0
	var a int
	// default to ""
	var str string
	// default to 0
	var b float32
	// default to false
	var flag bool

	fmt.Println(a, b, str)
	fmt.Printf("c: %v\n", flag)

	// inline declaration without type
	var temp interface{}
	temp = 1
	fmt.Printf("temp: %v\n", temp)
}
