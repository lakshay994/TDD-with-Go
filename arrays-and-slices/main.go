package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(Sum(numbers))
}

// Sum takes in Array of integers and return the sum of all numbers.
func Sum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}
