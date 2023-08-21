package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))       // Output: "fl"
	fmt.Println(longestCommonPrefix([]string{"Triangle", "Tricycle", "Tripod"})) // Output: "tri"
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	longestCommonPrefix := strs[0]

	for _, str := range strs[1:] {
		if len(str) > len(longestCommonPrefix) {
			str = str[:len(longestCommonPrefix)]
		}

		commonPrefix := ""
		for index, char := range str {
			if string(char) == string(longestCommonPrefix[index]) {
				commonPrefix += string(longestCommonPrefix[index])
			}
		}

		longestCommonPrefix = commonPrefix
	}

	return longestCommonPrefix
}
