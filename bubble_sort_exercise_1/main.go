package main

import "log"
import "strings"

func bubbleSort(list []string) []string {
	for i := 0; i < len(list); i++ {
		if !compare(list, i) {
			return list
		}
	}
	return list
}

func compare(list []string, index int) bool {
	firstIndex := 0
	secondIndex := 1
	numberOfItems := len(list)

	var firstName, secondName string
	isSwap := false

	for secondIndex < (numberOfItems - index) {
		firstName = list[firstIndex]
		secondName = list[secondIndex]

		if compareString(firstName, secondName) {
			list[firstIndex] = secondName
			list[secondIndex] = firstName
			isSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return isSwap
}

func compareString(name1, name2 string) bool {
	if strings.ToLower(name1) > strings.ToLower(name2) {
		return true
	} else {
		return false
	}
}

func main() {
	list := []string{"dog", "DOG", "cat", "Cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", 
	"porcupine", "dung beetle", "camel", 
	"steer", "bat", "BAt", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
	orderedList := bubbleSort(list)
	log.Println(orderedList)
}