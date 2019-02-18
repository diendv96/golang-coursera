package main

import "fmt"

func main() {
	// Function return multiple value, so the assignment must be equivalent
	//a1, a2 := foo(1, 2)
	x := 2
	//add(x)
	add2(&x)
	fmt.Println(x)

}

func foo(param1 int, param2 string) (int, string) {
	return param1, param2
}

// Pass arguments by value to function
func add(y int) {
	y = y + 1
}

func add2(y *int) {
	*y = *y + 1
}
