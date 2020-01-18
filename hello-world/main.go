package main

import "fmt"

const helloString = "Hello, "
const namasteString = "Namaste, "
const punjabiString = "Sat Sri Akal, "

func main() {
	fmt.Println(PrintString("World!", "punjabi"))
}

// PrintString - this is the domain. needs to be seperated from the side-effects
func PrintString(val, language string) string {
	if val == "" {
		val = "World"
	}
	return pickLanguage(language) + val
}

func pickLanguage(language string) (prefix string) {
	switch language {
	case "english":
		prefix = helloString
	case "hindi":
		prefix = namasteString
	case "punjabi":
		prefix = punjabiString
	}
	return
}