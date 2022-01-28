package main

import "fmt"

type person struct {
	firstName   string `json:"first_name,omitempty"`
	lastName    string `json:"last_name,omitempty"`
	contactInfo `json:"contact"`
}
type contactInfo struct {
	email string `json:"email,omitempty"`
	zip   int    `json:"zip,omitempty"`
}

func main() {
	p := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "jimbo@email.com",
			zip:   11111,
		},
	}
	s := []string{"hello", "world!"}
	fmt.Println(s)
	updateSlice(s)
	fmt.Println(s)
	p.print()
	p.updateName("Jim")
	p.print()
}
func (p *person) updateName(name string) {
	(*p).firstName = name
}
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateSlice(s []string) {
	s[0] = "test"
}
