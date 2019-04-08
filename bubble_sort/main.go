package main

import "log"

func bubbleSort(list []int) []int {
	var firstNumber, secondNumber int
	checkResult := 0
	for {
		for i := 0; i < len(list) - 1; i++ {
			// Đổi chỗ 2 số nếu số trước > số sau
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

func bubbleSort_2(list []int) []int {
	for i := 0; i < len(list); i++ {
		if !compare(list, i) {
			return list
		}
	}
	return list
}

func compare(list []int, index int) bool {
	firstIndex := 0
	secondIndex := 1
	numberOfItems := len(list)

	var firstNumber, secondNumber int
	isSwap := false

	for secondIndex < (numberOfItems - index) {
		firstNumber = list[firstIndex]
		secondNumber = list[secondIndex]

		if firstNumber > secondNumber {
			list[firstIndex] = secondNumber
			list[secondIndex] = firstNumber
			isSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return isSwap
}

func main() {
	list := []int{10,9,9,9,3,3}
	orderedList := bubbleSort_2(list)
	log.Println(orderedList)
}