// Take 2 strings s1 and s2 including only letters from a to z.
// Return a new sorted string (alphabetical ascending), the longest possible,
// containing distinct letters - each taken only once - coming from s1 or s2.
// 			::: example :::
//	a = "xyaabbbccccdefww"
//	b = "xxxxyyyyabklmopq"
//	longest(a, b) -> "abcdefklmopqwxy"

package main

import (
	"fmt"
	"sort"
	"strings"
)

func processString(a string, b string) string {
	concat, res, acc := a+b, []string{""}, ""

	for i := 0; i < len(concat); i++ {
		char := string(concat[i])

		if strings.Count(acc, char) < 1 {
			acc += string(concat[i])
			res = append(res, char)
		}
	}
	sort.Strings(res)

	return strings.Join(res, "")
}

func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	fmt.Println(processString(a, b)) // "abcdefklmopqwxy"
}
