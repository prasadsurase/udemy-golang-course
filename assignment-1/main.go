package main

import (
	"fmt"
)

func main() {
	data := []int{11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var output string

	for _, num := range data {
		if num%2 == 0 {
			output = "even"
		} else {
			output = "odd"
		}
		fmt.Println(num, "is", output)
	}
}
