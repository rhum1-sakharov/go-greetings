package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (p *person) updateName(newFirstname string) {
	p.firstname = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
