package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Input
	fmt.Print("Enter series of integer numbers: ")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numbers, err := parseInts(line)
	if err != nil {
		log.Fatal(err)
	}
	// Create a channel
	c := make(chan []int)

	// Divide
	numbersLen := len(numbers)
	divideSlice := numbersLen / 4
	index := 0
	for i := 0; i < 3; i++ {
		go sortSubArray(numbers[index : index+divideSlice], c)
		index = index + divideSlice
	}
	go sortSubArray(numbers[index:], c)

	sub1 := <- c
	sub2 := <- c
	sub3 := <- c
	sub4 := <- c

	mergeSub12 := merge(sub1, sub2)
	mergeSub34 := merge(sub3, sub4)
	mergeSub1234 := merge(mergeSub12, mergeSub34)
	fmt.Println("Sort result: ", mergeSub1234)
}

func parseInts(s string) ([]int, error) {
	var (
		fields  = strings.Fields(s)
		numbers = make([]int, len(fields))
		err     error
	)
	for i, f := range fields {
		numbers[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return numbers, nil
}

func sortSubArray(slice []int, c chan []int) []int {
	fmt.Println("Will sorting following subarray: ", slice)
	sort.Ints(slice)
	c <- slice
	return slice
}

func merge(left, right []int) []int {
	//make an []int of size of length of left + length of right
	result := make([]int, len(left)+len(right))

	for i := 0; len(left) > 0 || len(right) > 0; i++ {

		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}

	return result
}