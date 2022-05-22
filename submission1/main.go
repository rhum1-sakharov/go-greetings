package main

import "fmt"

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printIsOddOrEven(intSlice)
}

func printIsOddOrEven(intSlice []int) {
	for _, value := range intSlice {
		strIsOddOrEven := ""
		if value%2 == 0 {
			strIsOddOrEven = "odd"
		} else {
			strIsOddOrEven = "even"
		}
		fmt.Printf("%v => %v\n", value, strIsOddOrEven)
	}
}
