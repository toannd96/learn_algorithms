package main

import "log"

func bubbleSort(list []int) []int {
	var firstNumber, secondNumber int
	checkResult := 0
	for {
		for i := 0; i < len(list) - 1; i++ {
			firstNumber = list[i]
			secondNumber = list[i+1]
			if firstNumber > secondNumber {
				list[i] = secondNumber
				list[i+1] = firstNumber
				checkResult++
			}
		}

		if checkResult == 0 {
			return list;
		}

		checkResult = 0
	}
}

func main() {
	list := []int{4,1,9,9,3,3}
	orderedList := bubbleSort(list)
	log.Println(orderedList)
}