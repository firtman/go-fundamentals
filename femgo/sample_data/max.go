package sampledata

import (
	"time"

	"frontendmasters.com/goapp/data"
)

func GetData() []data.Printable {
	var courses []data.Printable

	max := data.Instructor{Name: "Maximiliano Firtman"}

	goCourse := data.Course{Name: "Go Fundamentals", Duration: 4}
	goCourse.Instructor = max

	kotlinWorkshop := data.Workshop{Date: time.Now().AddDate(0, 0, 15)}
	kotlinWorkshop.Name = "Kotlin for Android"
	kotlinWorkshop.Instructor = max

	swiftWorkshop, _ := data.NewWorkshop("Swift", 3, max, time.Now())

	return append(courses, goCourse, kotlinWorkshop, swiftWorkshop)
}
