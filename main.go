package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	Street string `json:"street" flag:"true"`
	Number int `json:"number"`
}

type Person struct {
	Name string `json:"name" flag:"true"`
	Age int `json:"age"`
	Address Address `json:"address"`
}

func main() {
	println("Hello, World!")

	// Create an example of person
	// person := Person{
	// 	Name: "John",
	// 	Age: 30,
	// 	Address: Address{
	// 		Street: "Main Street",
	// 		Number: 123,
	// 	},
	// }
	tryCastInterfaceToStruct()

}
type Person1 struct {
	Name string `json:"name" flag:"true"`
	Age int `json:"age"`
	LastName string `json:"last_name"`

}

type Person2 struct {
	Name string `json:"name" flag:"true"`
	Age int `json:"age"`
}

func tryCastInterfaceToStruct() {
	// Create a new struct
	person := Person1{
		Name: "John",
		Age: 30,
	}
	// person2 := Person2{
	// 	Name: "John",
	// 	Age: 30,
	// }
	// Create a new struct of the same type as the person struct
	// newPerson := Person{}

	// Convert the person struct to an interface
	personInterface := interface{}(person)
	// convert person to Person2
	newPerson, ok  := personInterface.(Person1)
	// Convert the interface back to a struct
	// newPerson, ok := personInterface.(Person)
	if !ok {
		fmt.Println("Failed to convert the interface to a struct")
	}

	fmt.Println(newPerson)

}

func addFlagToStruct(person Person) Person {
		newPerson := person

		// Get the type of the Person struct
		personType := reflect.TypeOf(person)
		// check if person is a pointer or struct
		if personType.Kind() == reflect.Ptr {
			personType = personType.Elem()
		}


		// Iterate over the fields of the Person struct
		for i := 0; i < personType.NumField(); i++ {
			field := personType.Field(i)

			// Check if the field has the "flag" tag set to true
			if field.Tag.Get("flag") == "true" {
				// Create a new field with the "_flag" postfix
				flagFieldName := field.Name + "_flag"
				flagFieldType := reflect.TypeOf(false)
				flagField := reflect.StructField{
					Name: flagFieldName,
					Type: flagFieldType,
					Tag:  reflect.StructTag(`json:"` + flagFieldName + `"`),
				}

				// Add the new field to the newPerson struct
				newPersonType := reflect.TypeOf(newPerson)
				fmt.Println(newPersonType)
				newPersonType = reflect.StructOf([]reflect.StructField{field, flagField})
				newPersonValue := reflect.New(newPersonType).Elem()
				newPersonValue.FieldByName(field.Name).Set(reflect.ValueOf(person).FieldByName(field.Name))
				newPersonValue.FieldByName(flagFieldName).Set(reflect.ValueOf(false))
				//The following line is not working: newPerson = newPersonValue.Interface().(Person)
				//fix it.
				newPerson = newPersonValue.Interface().(Person)
				//above line is throwing panic, fix it: interface conversion: interface {} is struct { Name string \"json:\\....
				
				



			}
		}

		return newPerson
	}


