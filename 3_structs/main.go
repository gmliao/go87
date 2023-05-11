package main

import "fmt"

type contectInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contectInfo
}

func (p *person) fullName() string {
	return (*p).firstName + " " + p.lastName
}

func (p *person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	alex := person{firstName: "Alex",
		lastName: "Anderson",
		age:      20,
		contectInfo: contectInfo{
			email:   "alex@abc.com",
			zipCode: 123456,
		},
	}

	fmt.Println(alex.fullName())
	alex.age = 45
	alex.updateName("guanming")
	alex.print()
}
