package anagram

import "testing"

var cases = []struct {
	s1       string
	s2       string
	expected bool
}{
	{"", "", false},
	{"a", "a", false},
	{"foobar", "foobar", false},
	{"foobar", "foo", false},
	{"foobar", "barqux", false},
	{"file", "life", true},
	{"anagram", "nagaram", true},
}

func TestAreAnagramsOne(t *testing.T) {
	for _, c := range cases {
		if got := areAnagramsOne(c.s1, c.s2); got != c.expected {
			t.Errorf("areAnagrams(%q, %q) = %v", c.s1, c.s2, got)
		}
	}
}

func TestAreAnagramsTwo(t *testing.T) {
	for _, c := range cases {
		if got := areAnagramsTwo(c.s1, c.s2); got != c.expected {
			t.Errorf("areAnagrams(%q, %q) = %v", c.s1, c.s2, got)
		}
	}
}
