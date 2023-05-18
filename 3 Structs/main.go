package main

import "fmt"

type contectInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	age       int
	contectInfo
}

type Student struct {
	Person // embedded type
	grade  int
}

type Housewife struct {
	Person  // embedded type
	husband Person
}

func (p *Person) fullName() string {
	return (*p).firstName + " " + p.lastName
}

func (p *Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	alex := Person{firstName: "Alex",
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

	jimmy := Student{Person: Person{firstName: "Jimmy", lastName: "Lai"}, grade: 100}
	fmt.Println(jimmy.fullName())
	fmt.Println(jimmy.firstName)
	fmt.Printf("%+v\n", jimmy)

	// housewife
	housewife := Housewife{Person: Person{firstName: "Mary", lastName: "Lee"}, husband: jimmy.Person}

	fmt.Println(housewife.fullName())
	fmt.Println(housewife.husband.fullName())
}
