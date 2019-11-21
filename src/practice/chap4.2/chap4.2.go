package main

import "fmt"

func main() {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	fmt.Println("irregular:", irregularMatrix)
	slice := Flatten(irregularMatrix)
	fmt.Printf("shape(1x%d): %v\n", len(slice), slice)
}

func Flatten(inputInts [][]int) []int {
	outputIns := []int{}
	for _, line := range inputInts {
		outputIns = append(outputIns, line...)
	}
	return outputIns

}
