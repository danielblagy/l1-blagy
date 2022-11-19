package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name         string
	PlaceOfBirth string
	Birthday     time.Time
}

func (h Human) getAge() int {
	currentDate := time.Now()
	yearsElapsed := currentDate.Year() - h.Birthday.Year()
	if currentDate.YearDay() < h.Birthday.YearDay() {
		yearsElapsed -= 1
	}
	return yearsElapsed
}

type Action struct {
	// embedding Human struct into Action
	Human
	Title string
}

func (a Action) work() {
	fmt.Println(a.Name, "(", a.Title, ") is working")
}

func main() {
	myBirthday, _ := time.Parse("2006-01-02", "2000-12-13")
	me := Action{
		Human{
			Name:         "Daniel Blagy",
			PlaceOfBirth: "Sevastopol, Ukraine",
			Birthday:     myBirthday,
		},
		"Intern",
	}
	// Action struct "inherits" Human fields and methods
	fmt.Println("Name:", me.Name)
	fmt.Println("Place of birth:", me.PlaceOfBirth)
	fmt.Println("Birthday:", me.Birthday.Format("2006-01-02"))
	fmt.Println("Age:", me.getAge())

	me.work()
}
