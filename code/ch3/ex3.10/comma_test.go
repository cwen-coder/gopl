package comma

import "testing"

var cases = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"0", "0"},
	{"1", "1"},
	{"123", "123"},
	{"1234", "1,234"},
	{"12345", "12,345"},
	{"123456789", "123,456,789"},
	{"1234567890", "1,234,567,890"},
}

func TestComma(t *testing.T) {
	for _, c := range cases {
		if got := comma(c.input); got != c.expected {
			t.Errorf("comma(%q) = %q", c.input, got)
		}
	}
}
