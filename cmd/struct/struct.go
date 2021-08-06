package main

import (
	"fmt"
)

type Student struct {
	fname string
	lname string

	ageStrFn func(a int) string
}

func (s *Student) GetMyName() string {
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}

func MyAge(a int) string {
	return fmt.Sprintf("I'm %d year old", a)
}

func main() {
	tinnt := Student{
		fname:    "Tin",
		lname:    "Nguyen",
		ageStrFn: MyAge, // it will be a pointer
	}

	fmt.Println(tinnt.ageStrFn(30))

	fmt.Println(tinnt)
}
