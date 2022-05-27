package main

import (
	"fmt"
	"pb_test/pb/person"
)

func main() {
	var p person.Person
	one := p.TestOneOf.(*person.Person_One)
	one.One = "one"

	switch p.TestOneOf.(type) {
	case *person.Person_One:
		fmt.Println("one")
	case *person.Person_Two:
		fmt.Println("two")
	case *person.Person_Three:
		fmt.Println("three")
	}
}
