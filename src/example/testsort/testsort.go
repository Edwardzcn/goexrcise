package main

import (
	"fmt"
	"sort"
	"strings"
)

type FoldedString []string

func (slice FoldedString) Len() int {
	return len(slice)
}

func (slice FoldedString) Less(i int, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedString) Swap(i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	test1 := make(FoldedString, 2)
	test1 = append(test1, "who")
	test1 = append(test1, "What")
	test1 = append(test1, "Answer")
	fmt.Println(test1)
	sort.Sort(test1)
	fmt.Println(test1)
}
