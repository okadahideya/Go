package main

import (
	"fmt"
)

func main() {
	var s string = "Hellow Go"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	main_array()
}

func main_array() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	arr4 := [...]string{"A", "B"}
	fmt.Println(arr4)

	var x interface{}
	fmt.Println(x)
}
