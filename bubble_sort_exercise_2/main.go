package main

import "log"

type Person struct {
	FirstName string
	LastName string
}

func bubbleSort(list []Person) []Person {
	for i := 0; i < len(list); i++ {
		if !compareLastName(list, i) {
			break;
		}
	}

	for i := 0; i < len(list); i++ {
		if !compareFirstName(list, i) {
			return list
		}
	}

	return list
}

func compareLastName(list []Person, index int) bool {
	firstIndex := 0
	secondIndex := 1
	numberOfItems := len(list)

	var firstName, secondName string
	var firstPerson, secondPerson Person
	isSwap := false

	for secondIndex < (numberOfItems - index) {
		firstPerson = list[firstIndex]
		secondPerson = list[secondIndex]
		
		firstName = firstPerson.LastName
		secondName = secondPerson.LastName

		if firstName > secondName {
			list[firstIndex] = secondPerson
			list[secondIndex] = firstPerson
			isSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return isSwap
}

func compareFirstName(list []Person, index int) bool {
	firstIndex := 0
	secondIndex := 1
	numberOfItems := len(list)

	var firstName, secondName string
	var firstPerson, secondPerson Person
	isSwap := false

	for secondIndex < (numberOfItems - index) {
		firstPerson = list[firstIndex]
		secondPerson = list[secondIndex]
		
		firstName = firstPerson.FirstName
		secondName = secondPerson.FirstName

		if firstName < secondName && firstPerson.LastName > secondPerson.LastName {
			list[firstIndex] = secondPerson
			list[secondIndex] = firstPerson
			isSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return isSwap
}

func main() {
	list := []Person{
		{
			FirstName: "Jon",
			LastName: "Calhoun",
		},
		{
			FirstName: "Bob",
			LastName: "Smith",
		},
		{
			FirstName: "John",
			LastName: "Smith",
		},
		{
			FirstName: "Larry",
			LastName: "Smith",
		},
		{
			FirstName: "Joseph",
			LastName: "Costanza",
		},
		{
			FirstName: "George",
			LastName: "Costanza",
		},
		{
			FirstName: "Jerry",
			LastName: "Costanza",
		},
		{
			FirstName: "Shane",
			LastName: "Calhoun",
		},
	}
	orderedList := bubbleSort(list)
	log.Println(orderedList)
}