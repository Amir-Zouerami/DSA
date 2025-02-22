package main

import (
	"fmt"

	leetcode "github.com/Amir-Zouerami/DSA/problems"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	output := leetcode.LongestConsecutive(nums)

	fmt.Println("output:", output)
}
