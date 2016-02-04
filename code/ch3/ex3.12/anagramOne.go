package anagram

import (
	"strings"
)

func areAnagramsOne(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}
	for _, v := range s1 {
		index := strings.Index(s2, string(v))
		if index == -1 {
			return false
		}
		if index+1 >= len(s2) {
			s2 = s2[:index]
		} else {
			s2 = s2[:index] + s2[index+1:]
		}
	}
	return true
}
