package comma

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := n - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
