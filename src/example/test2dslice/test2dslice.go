package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// debug参数设定为 ["3","4","5"]
	if len(os.Args) < 4 {
		fmt.Println("This program needs three int:")
		os.Exit(1)
	}
	grid1 := make([][]int, 3)
	for index := 0; index < 3; index++ {
		parint, _ := strconv.ParseInt(os.Args[index+1], 10, 0)
		grid1[index] = make([]int, parint)
	}
	fmt.Println(grid1)
	grid1[0][0], grid1[1][1], grid1[2][2] = 1, 2, 3
	fmt.Println(grid1)
	// [[0 0 0] [0 0 0 0] [0 0 0 0 0]]
	// [[1 0 0] [0 2 0 0] [0 0 3 0 0]]
	// 符合预期

	grid2 := make([]int, 5)
	for _, i := range grid2 {
		fmt.Println(i)
		i = 2
	}
	fmt.Println(grid2)
	for k, _ := range grid2 {
		grid2[k] = 2
	}
	fmt.Println(grid2)
	grid3 := make([]int, 0, 10)
	fmt.Println(grid3)
	grid3 = grid3[:10]
	fmt.Println(grid3)
}
