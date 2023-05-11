package data

import "fmt"

type Duration float32

type Course struct {
	Id          int
	Name        string
	Slug        string
	Description string
	Legacy      bool
	Duration    Duration
	Instructor  Instructor
}

func (c Course) Print() string {
	return fmt.Sprintf("%v (%.2f hours) by %v", c.Name, c.Duration, c.Instructor.Name)
}
