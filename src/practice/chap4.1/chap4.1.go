package main

import "fmt"

func main() {
	slice := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	fmt.Println("Original:", slice)
	slice = UniqueInts(slice)
	fmt.Println("Unique:  ", slice)
}

func UniqueInts(inputInts []int) []int {
	var outputInts []int
	var cntMap = make(map[int]int, len(inputInts))
	for _, value := range inputInts {
		if _, find := cntMap[value]; find == false {
			cntMap[value] = 1
			outputInts = append(outputInts, value)
		}
	}
	return outputInts
}
