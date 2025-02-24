package main

import (
	"fmt"

	leetcode "github.com/Amir-Zouerami/DSA/problems"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}

	output := leetcode.EvalRPN(tokens)

	fmt.Println("output:", output)
}
