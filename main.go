package main

import (
	"fmt"

	leetcode "github.com/Amir-Zouerami/DSA/problems"
)

func main() {
	strs := []string{"neet", "code", "love", "you"}

	encoded := leetcode.Encode(strs)
	decoded := leetcode.Decode(encoded)

	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)
}
