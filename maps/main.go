package main

import "errors"

// ErrorResponse is returned if a searched word is not found inside the dictionary
var ErrorResponse string = "could not find the word you are looking for"

// Dictionary is a struct that maps a sting key to a string value
type Dictionary map[string]string

// Search takes in a map of string to string and a word and returns the value of that word from the dictionary
func (d *Dictionary) Search(dictionary map[string]string, word string) (string, error) {
	if value, ok := dictionary[word]; ok {
		return value, nil
	}
	return "", errors.New(ErrorResponse)
}
