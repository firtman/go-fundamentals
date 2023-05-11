package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a string
	str := "Hello, World!"

	// Get the length of the string
	length := len(str)
	fmt.Println("The length of the string is", length)

	// Convert the string to uppercase
	upperCaseStr := strings.ToUpper(str)
	fmt.Println("The uppercase string is", upperCaseStr)

	// Convert the string to lowercase
	lowerCaseStr := strings.ToLower(str)
	fmt.Println("The lowercase string is", lowerCaseStr)

	// Replace a substring in the string
	replacedStr := strings.Replace(str, "World", "Universe", 1)
	fmt.Println("The replaced string is", replacedStr)

	// Split the string into a slice of strings
	slice := strings.Split(str, " ")
	fmt.Println("The slice of strings is", slice)

	// Join a slice of strings into a string
	joinedStr := strings.Join(slice, " ")
	fmt.Println("The joined string is", joinedStr)
}
