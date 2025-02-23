package main

import (
	leetcode "github.com/Amir-Zouerami/DSA/problems"
)

func main() {
	minStack := leetcode.Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	minStack.GetMin()
	minStack.Pop()
	minStack.Top()
	minStack.GetMin()

	// output := leetcode.LongestConsecutive(nums)

	// fmt.Println("output:", output)
}
