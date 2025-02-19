package leetcode

import (
	"strconv"
	"strings"
)

func Encode(strs []string) string {
	var encodedStr strings.Builder

	for _, str := range strs {
		encodedStr.WriteString(strconv.Itoa(len(str)) + "#" + str)
	}

	return encodedStr.String()
}

func Decode(encoded string) []string {
	i := 0
	decodedArray := []string{}

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		runes := []rune(encoded)

		length, _ := strconv.Atoi(string(runes[i:j]))
		item := string(runes[j+1 : j+1+length])

		decodedArray = append(decodedArray, item)

		i = j + 1 + length
	}

	return decodedArray
}
