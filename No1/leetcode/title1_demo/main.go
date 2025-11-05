package main

import (
	"Golang_study/No1/leetcode/title1"
	"fmt"
)

func main() {
	// 两数之和 只会存在一个答案
	var nums []int = []int{2, 7, 11, 15}
	fmt.Printf("两数之和的下标为:%v \n", title1.Twosum(nums, 9))

	//合并区间
	var intervals [][]int = [][]int{{1, 2}, {3, 4}, {1, 6}, {7, 8}, {7, 10}}
	fmt.Printf("合并重叠区间最终结果为: %v \n", title1.Merge(intervals))

	// 删除有序数组中的重复项 双指针
	var fastSlowDuplicates []int = []int{0, 0, 1, 1, 2, 2, 2, 3, 4, 5, 5, 6}
	fmt.Printf("删除有序数组中的重复项后还剩:%d 个数 \n", title1.RemoveDuplicates(fastSlowDuplicates))

	// 加一
	var plusOnes []int = []int{9}
	fmt.Printf("plusOnes加一后: %d \n", title1.PlusOne(plusOnes))

	// 有效的括号
	var validBrackets string = "([{}])"
	fmt.Printf("判断括号: %s ,是否有效:%t \n", validBrackets, title1.IsValid(validBrackets))

	// 判断回文数
	var palindrome int = 12321
	fmt.Printf("判断【%d】是否为回文数: %t \n", palindrome, title1.IsPalindrome(palindrome))

	// 只出现一次的数字
	var numbers []int = []int{2, 2, 1}
	fmt.Printf("只出现一次的数字: %d\n", title1.SingleNumber(numbers))

	// 最长公共前缀
	var strs []string = []string{"flower", "flow", "flight"}
	fmt.Printf("最长公共前缀为: %s\n", title1.LongestCommonPrefix(strs))
}
