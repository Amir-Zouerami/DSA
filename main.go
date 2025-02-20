package main

import (
	"fmt"

	leetcode "github.com/Amir-Zouerami/DSA/problems"
)

func main() {
	nums := []int{1, 2, 3, 4}

	output := leetcode.ProductExceptSelf(nums)

	fmt.Println("output:", output)
}
