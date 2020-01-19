package main

import "fmt"

func main() {
	fmt.Println(Add(11, 12))
}

// Add take in two integers and adds them.
func Add(first, second int) (sum int) {
	return first + second
}
