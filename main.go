package main

import (
	"fmt"

	leetcode "github.com/Amir-Zouerami/DSA/problems"
)

func main() {
	data := leetcode.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(data)
}
