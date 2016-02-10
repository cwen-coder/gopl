package spaces

import (
	"bytes"
	"testing"
)

var cases = []struct {
	input    []byte
	expected []byte
}{
	{[]byte(""), []byte("")},
	{[]byte("a"), []byte("a")},
	{[]byte("foo bar"), []byte("foo bar")},
	{[]byte("foo\tbar"), []byte("foo bar")},
	{[]byte("\tfoo\t\vbar\r\n"), []byte(" foo bar ")},
}

func TestTrimSpaces(t *testing.T) {
	for _, c := range cases {
		if got := trimSpaces(c.input); !bytes.Equal(got, c.expected) {
			t.Errorf("trimSpaces(%v) = %v", c.input, got)
		}
	}
}
