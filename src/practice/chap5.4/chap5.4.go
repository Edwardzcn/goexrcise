package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	testData := [][]string{
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/goeg/prefix/extra"},
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/prefix/extra"},
		{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
			"/pecan/π/prefix/extra"},
		{"/pecan/π/circle", "/pecan/π/circle/prefix",
			"/pecan/π/circle/prefix/extra"},
		{"/home/user/goeg", "/home/users/goeg",
			"/home/userspace/goeg"},
		{"/home/user/goeg", "/tmp/user", "/var/log"},
		{"/home/mark/goeg", "/home/user/goeg"},
		{"home/user/goeg", "/tmp/user", "/var/log"},
	}
	for _, data := range testData {
		fmt.Printf("[")
		gap := ""
		for _, datum := range data {
			fmt.Printf("%s\"%s\"", gap, datum)
			gap = " "
		}
		fmt.Println("]")
		cp := CommonPrefix(data)
		cpp := CommonPathPrefix(data)
		equal := "=="
		if cpp != cp {
			equal = "!="
		}
		fmt.Printf("char ⨉ path prefix: \"%s\" %s \"%s\"\n\n",
			cp, equal, cpp)
	}
}

func CommonPrefix(input []string) string {
	runeSlice := make([][]rune, len(input))
	for i, s := range input {
		runeSlice[i] = []rune(s)
	}
	if len(runeSlice) == 0 || len(runeSlice[0]) == 0 {
		return ""
	}
	// 横行扫描方法
	// 转化为rune二维数组
	// length := len(runeSlice[0])
	// for i, line := range runeSlice {
	// 	if i == 0 {
	// 		continue
	// 	} else {
	// 		for j := 0; j < length; j++ {
	// 			if line[j] != runeSlice[0][j] {
	// 				length = j
	// 				break
	// 			}
	// 		}
	// 		if length == 0 {
	// 			return ""
	// 		}
	// 	}
	// }
	// return string(runeSlice[0][:length])
	// 纵列扫描法
	ans := make([]rune, 0, len(runeSlice[0]))
	for length := 0; length < len(runeSlice[0]); length++ {
		for row := 1; row < len(runeSlice); row++ {
			if length >= len(runeSlice[row]) || runeSlice[row][length] != runeSlice[0][length] {
				return string(ans)
			}
		}
		ans = append(ans, runeSlice[0][length])
	}
	return string(ans)
}

func CommonPathPrefix(input []string) string {
	const separator = string(filepath.Separator)
	contents := make([][]string, len(input))
	for i, line := range input {
		contents[i] = strings.Split(line, separator)
		if strings.HasPrefix(line, separator) {
			contents[i] = append([]string{separator}, contents[i]...)
		}
	}
	if len(contents) == 0 || len(contents[0]) == 0 {
		return ""
	}
	// 纵列扫描
	ans := make([]string, 0, len(contents[0]))
FINISH:
	for col := 0; col < len(contents[0]); col++ {
		word := contents[0][col]
		for row := 1; row < len(contents); row++ {
			if col >= len(contents[row]) || word != contents[row][col] {
				break FINISH
			}
		}
		ans = append(ans, word)
	}
	// filepath.Join()
	// Join joins any number of path elements into a single path, adding a Separator if necessary.
	// 注意这个if necessary很是微妙
	return filepath.Join(ans...)

}
