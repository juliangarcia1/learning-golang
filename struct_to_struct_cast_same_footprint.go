package main

import (
	"fmt"
	"reflect"
)

type Person1 struct {
	Name string `json:"name" flag:"true"`
	Age int `json:"age"`
	// LastName string `json:"last_name"`

}

type Person2 struct {
	Name string `json:"name" flag:"true"`
	Age int `json:"age"`
}

func CastInterfaceToStruct() {
	//------------- INPUT
	// Create a new struct
	person := Person1{
		Name: "John",
		Age: 30,
	}

	//------------- PROCESS
	// Convert the person struct to an interface
	// personInterface := interface{}(person)
	// convert person to Person2. Convert the interface back to a struct
	newPerson := Person2(person)
	// newPerson, ok := personInterface.(Person)
	// if !ok {
	// 	fmt.Println("Failed to convert the interface to a struct")
	// }
	
	//------------- OUTPUT
	fmt.Println("person=", person, 
				"\nOriginal Type:", reflect.ValueOf(person).Type(),
				"\nnewPerson:", newPerson,
				"\nnewPerson Type:", reflect.ValueOf(newPerson).Type(),)

}