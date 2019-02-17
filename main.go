package main

import "fmt"

type Person struct {
	name string
	addr string
	age  int
}

func main() {
	//trunc()
	//findian()
	var x = [3]int{1, 2, 3}
	for i, v := range x {
		fmt.Println("ind %d, val %d", i, v)
	}
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(s1)
	fmt.Println(s2)

	sli := make([]int, 10)
	fmt.Println(sli)

	sli2 := make([]int, 10, 15)
	sli2 = append(sli2, 9999)
	fmt.Println(sli2)

	idMap := map[string]int{
		"Joe":   123,
		"Craig": 999,
	}
	idMap["Bob"] = 456
	fmt.Println(idMap["Craig"])
	fmt.Println(idMap["Bob"])

	delete(idMap, "Joe")

	fmt.Println(idMap)

	for key, val := range idMap {
		fmt.Println(key, ":", val)
	}

	var p1 Person
	p1.name = "Craig"
	p1.addr = "Hanoi"
	p1.age = 22

	p2 := new(Person)
	p2.name = "Test"

	fmt.Println(p1)
}
