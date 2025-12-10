package Problems

import (
	"fmt"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	runeLocationMap := make(map[rune]int)
	maxLength := 0
	// currentLength:=1;
	uniqueStringBeginner := 0
	for index, char := range s {
		indexOfCurrentChar, ok := runeLocationMap[char]
		if ok {
			if uniqueStringBeginner < (indexOfCurrentChar + 1) {
				uniqueStringBeginner = indexOfCurrentChar + 1
			}
		}
		currentLength := (index - uniqueStringBeginner) + 1
		fmt.Printf("Length %d uniqueStringBeginner %d index %d rune %c", currentLength, uniqueStringBeginner, index, char)
		fmt.Println("")
		if currentLength > maxLength {
			maxLength = currentLength
		}
		runeLocationMap[char] = index
	}
	return maxLength
}

func LengthOfLongestSubstringProblem() {
	fmt.Println("--------------------------------")
	fmt.Println("Length of longest substring")
	fmt.Println("Description:")
	fmt.Println(`Given a string s, find the length of the longest substring without duplicate characters.`)
	fmt.Println("Example 1:")
	fmt.Println("Input: s = abcabcbb")
	fmt.Println("Output: ")
	start := time.Now()
	lengthOfLongestSubstring("abcabcbb")
	fmt.Printf("Time taken: %v\n", time.Since(start))
	fmt.Println("Explanation: The answer is abc, with the length of 3.")
	fmt.Println("Example 2:")
	fmt.Println("Input: s = bbbbb")
	fmt.Println("Output: ")
	start = time.Now()
	lengthOfLongestSubstring("bbbbb")
	fmt.Printf("Time taken: %v\n", time.Since(start))
	fmt.Println("Explanation: The answer is b, with the length of 1.")
	fmt.Println("Example 3:")
	fmt.Println("Input: s = pwwkew")
	fmt.Println("Output: ")
	start = time.Now()
	lengthOfLongestSubstring("pwwkew")
	fmt.Printf("Time taken: %v\n", time.Since(start))
	fmt.Println("Explanation: The answer is wke, with the length of 3.")
	fmt.Println("Notice that the answer must be a substring, wke is a subsequence and not a substring.")
	fmt.Println("--------------------------------")
	fmt.Println("")
}
