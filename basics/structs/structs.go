package main

import "fmt"

func main() {
	// make a person
	marlon := Person{
		name: "marlon",
		address: Address{streetName: "matson st", streetNumber: 43}}
	marlon.address.toString()
}

// address struct
type Address struct {
	streetName   string
	streetNumber int
}

// person struct
type Person struct {
	name    string
	address Address
}

// function: print address
func (this Address) toString() {
	fmt.Printf("street name: %v\nno.: %v", this.streetName, this.streetNumber)
}
