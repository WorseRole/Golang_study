package main

import (
	"fmt"

	"Golang_study/No1/leetcode/title1"
	_ "Golang_study/No1/pgk1"
	_ "Golang_study/No1/pgk2"
)

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
	fmt.Println("main init method invoked")
}

func main() {

	// 最长公共前缀
	var strs []string = []string{"flower", "flow", "flight"}
	fmt.Println(title1.LongestCommonPrefix(strs))

	fmt.Println("main method invoked!")
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}
