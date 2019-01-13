package main

import "fmt"

func trunc() {
	var x float32
	fmt.Println("Enter a float number:")
	fmt.Scan(&x)
	fmt.Println("The truncated number you just entered is ", int32(x))
}
