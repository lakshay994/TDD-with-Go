package main

import "fmt"

func main() {
	fmt.Println(Iterate("a", 10))
}

// Iterate takes in a string(val) and a integer(repeatCount) and repeats the string repeatCount number of times.
func Iterate(val string, repeatCount int) (res string) {
	for i := 0; i < repeatCount; i++ {
		res += val
	}
	return
}
