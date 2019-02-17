package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	intList := make([]int, 3)

	var scanValue string
	var index int = 0
	for {
		fmt.Print("Enter an integer", ": ")
		_, err := fmt.Scan(&scanValue)
		exitCode := strings.ToLower(scanValue)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if exitCode == "x" {
			sort.Ints(intList)
			fmt.Println(intList)
			os.Exit(0)
		}
		parseIntScanValue, err := strconv.Atoi(scanValue)
		if index < 3 {
			intList[index] = parseIntScanValue
			index++
		} else {
			intList = append(intList, parseIntScanValue)
		}
	}
}
