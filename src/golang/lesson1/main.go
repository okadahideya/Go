package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Hellow World")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)

	var s string = "Hellow go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t,f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2,s2)

	i4 := 400 
	fmt.Println(i4)
}