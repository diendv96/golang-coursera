package main

import "fmt"

type P struct {
	x string
	y int
}

func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}
