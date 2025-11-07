package main

import (
	"fmt"

	"Golang_study/No1/leetcode/title1"
)

func init() {
	fmt.Println("main init method invoked")
}

func main() {

	// 最长公共前缀
	var strs []string = []string{"flower", "flow", "flight"}
	fmt.Println(title1.LongestCommonPrefix(strs))

	fmt.Println("main method invoked!")
}
