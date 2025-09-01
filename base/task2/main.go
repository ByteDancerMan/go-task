package main

import (
	"strconv"

	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//例如，121 是回文，而 123 不是

func isPalindrome(x int) bool { 
	// 转换为字符串
	stringX := strconv.Itoa(x)

	// 创建一个数组
	digits := make([]int, len(stringX))
	
	// 将字符串转为数组
	for i, char := range stringX {
		// 将每个字符转换为对应的数字并存储在切片中
		digits[i] = int(char - '0')
	}


	for i,j := 0, len(digits) -1; i<j; i,j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

func main() { 
	isPalindrome(121)
}