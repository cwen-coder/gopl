package anagram

import "sort"

type sortableRunes []rune

func (sr sortableRunes) Len() int {
	return len(sr)
}

func (sr sortableRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr sortableRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func sortChars(s string) string {
	sorted := sortableRunes(s)
	sort.Sort(sorted)
	return string(sorted)
}

func areAnagramsTwo(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}
	if sortChars(s1) == sortChars(s2) {
		return true
	} else {
		return false
	}
}
