package main

import "fmt"

func main() {
	type Family struct {
		Name         string
		Relationship string
		DOB          string
	}
	person1 := Family{
		Name:         "Lukcas",
		Relationship: "Father",
		DOB:          "28-9-1964",
	}
	person2 := Family{
		Name:         "Zeline",
		Relationship: "Daughter",
		DOB:          "2-12-1990",
	}
	person3 := Family{
		Name:         "Maryan",
		Relationship: "Mother",
		DOB:          "18-2-1969",
	}
	person4 := Family{
		Name:         "Kenton",
		Relationship: "Son",
		DOB:          "5-5-1997",
	}

	familymembers := []Family{
		person1,
		person2,
		person3,
		person4,
	}
	for _, members := range familymembers {
		fmt.Printf("%s ,%s \n", members.name, members.DOB)
	}
}
