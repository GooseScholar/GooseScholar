package main

import (
	"fmt"
	"strings"
)

func main() {

	// Everyone knows passphrases. One can choose passphrases from poems, songs, movies names and
	// so on but frequently they can be guessed due to common cultural references. You can get your
	// passphrases stronger by different means. One is the following:

	// choose a text in capital letters including or not digits and non alphabetic characters,

	// 1. shift each letter by a given number but the transformed letter must be a letter (circular
	// shift),
	// 2. replace each digit by its complement to 9,
	// 3. keep such as non alphabetic and non digit characters,
	// 4. downcase each letter in odd position, upcase each letter in even position (the first
	// 	character is in position 0),
	// 5. reverse the whole result.

	// Example:

	// your text: "BORN IN 2015!", shift 1

	// 1 + 2 + 3 -> "CPSO JO 7984!"

	// 4 "CpSo jO 7984!"

	// 5 "!4897 Oj oSpC"

	// With longer passphrases it's better to have a small and easy program. Would you write it?

	// https://en.wikipedia.org/wiki/Passphrase

	fmt.Println(PlayPass("I`{ 69 LOVE 24 YOU ZZZZ 0 1 2 3 4 5 6 7 8 9 !!!", 3))
}

func PlayPass(s string, n int) string {
	if n >= 0 {
		s = strings.ToLower(s)
		a := []byte(s)
		var answer string

		for i := 0; i < len(s); i++ {
			if a[i] >= 97 && a[i] <= 122 {
				for j := 0; j < n; j++ {
					if a[i] == 122 {
						a[i] = 97
					} else {
						a[i]++
					}
				}
			} else if a[i] >= 48 && a[i] <= 57 {
				a[i] = 105 - a[i]
			}
		}

		for i := 1; i <= len(s); i++ {
			if len(s)%2 == 0 {
				if i%2 == 0 {
					answer += strings.ToUpper(string(a[len(s)-i]))
				} else {
					answer += string(a[len(s)-i])
				}
			} else {
				if i%2 == 1 {
					answer += strings.ToUpper(string(a[len(s)-i]))
				} else {
					answer += string(a[len(s)-i])
				}
			}
		}
		return answer
	} else {
		return "error"
	}
}
