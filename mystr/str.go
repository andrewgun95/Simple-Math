// Package mystr is contains all function relate to string manipulation
package mystr

import (
	"strings"
)

// Join will join all words together with space separator
func Join(words []string) string {
	return strings.Join(words, " ")
}

// Cat will concat all words together with space separator
func Cat(words []string) string {
	s := ""
	for i:=0; i < len(words);i++ {
		s += words[i]
		if i < len(words) - 1 {
			s += " "
		}
	}
	return s
}