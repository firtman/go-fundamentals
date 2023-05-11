package main

import (
	data "frontendmasters.com/goapp/data"
	sampledata "frontendmasters.com/goapp/sample_data"
)

var courses []data.Printable

func main() {
	courses = sampledata.GetData()
	for _, c := range courses {
		println(c.Print())
	}
}
