package main

import "fmt"

const helloString = "Hello "

func main() {
	fmt.Println(HelloString("World!"))
}

// HelloString - this is the domain. needs to be seperated from the side-effects
func HelloString(val string) string {
	if val == "" {
		val = "World"
	}
	return helloString + val
}
