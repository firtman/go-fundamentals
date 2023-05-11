package data

import (
	"fmt"
	"time"
)

type Workshop struct {
	Course
	Date         time.Time
	MaxAttendees int
}

func NewWorkshop(name string, duration Duration, instructor Instructor,
	date time.Time) (*Workshop, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name expected. %v characters is an invalid name", len(name))
	} else {
		aux := Workshop{}
		aux.Name = name
		aux.Duration = duration
		aux.Instructor = instructor
		aux.Date = date
		return &aux, nil
	}
}
