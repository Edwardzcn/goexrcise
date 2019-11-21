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
	fmt.Printf("1x%d: %v\n", len(slice), slice)
	fmt.Printf(" 3x%d: %v\n", CountRows(slice, 3), Make2D(slice, 3))
	fmt.Printf(" 4x%d: %v\n", CountRows(slice, 4), Make2D(slice, 4))
	fmt.Printf(" 5x%d: %v\n", CountRows(slice, 5), Make2D(slice, 5))
	fmt.Printf(" 6x%d: %v\n", CountRows(slice, 6), Make2D(slice, 6))
}

func Flatten(inputInts [][]int) []int {
	outputIns := []int{}
	for _, line := range inputInts {
		outputIns = append(outputIns, line...)
	}
	return outputIns

}

func Make2D(inputInts []int, column int) [][]int {
	matrix := make([][]int, CountRows(inputInts, column))
	for i, x := range inputInts {
		row := i / column
		col := i % column
		if matrix[row] == nil {
			matrix[row] = make([]int, column)
		}
		matrix[row][col] = x
	}
	return matrix
}

func CountRows(inputInts []int, column int) int {
	length := len(inputInts)
	if length%column == 0 {
		return length / column
	} else {
		return length/column + 1
	}
}
