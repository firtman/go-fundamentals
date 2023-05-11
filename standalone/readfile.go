package main

import (
	"io/ioutil"
	"os"
)

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func writeFile(filename string, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	directory, _ := os.Getwd()
	println(readFile(directory + "/first.go"))
	writeFile(directory+"/test.txt", "I'm a new file!")
}
