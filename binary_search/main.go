package main

import "fmt"

func binarySearch(list []int, number int) int {
	firstIndex := 0
	lastIndex := len(list) - 1

	for firstIndex <= lastIndex {
		mediumIndex := (firstIndex + lastIndex)/2
		if number == list[mediumIndex] {
			return mediumIndex
		} else if number > list[mediumIndex] {
			firstIndex = mediumIndex + 1
		} else {
			lastIndex = mediumIndex - 1
		}
	}

	return -1
}

func main(){
	list := []int{1,2,4,6,19,20,82,109,453,567,990,1001}
	indexFound := binarySearch(list, 453)
	if indexFound != -1 {
		fmt.Printf("Số cần tìm nằm ở vị trí %d \n", indexFound)
		fmt.Printf("Số cần tìm là: %d", list[indexFound])
	} else {
		fmt.Println("Không tìm thấy")
	}
}