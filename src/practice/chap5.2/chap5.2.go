// 非递归的回文判断判断函数

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Holds a reference to a function
//  创建一个引用后面用来指向ASCII或者UTF8两种回文判断函数
var IsPalindrome func(string) bool

func init() {
	if os.Args[1] == "-a" || os.Args[1] == "--ascii" {
		os.Args = append(os.Args[:1], os.Args[2:]...)
		IsPalindrome = func(input string) bool {
			length := len(input)
			for i := 0; i <= length/2; i++ {
				if input[i] != input[length-1-i] {
					return false
				}
			}
			return true
		}
	} else {
		IsPalindrome = func(input string) bool {
			runes := []rune(input)
			length := len(runes)
			// fmt.Println(length)
			for i := 0; i <= length/2; i++ {
				// fmt.Printf("--%d---%d--\n", runes[i], runes[length-i-1])
				if runes[i] != runes[length-1-i] {
					return false
				}
			}
			return true

		}
	}

}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s [-a|--ascii] word1 [word2 [... wordN]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	words := os.Args[1:]
	for _, word := range words {
		fmt.Printf("%5t %q\n", IsPalindrome(word), word)
	}
}
