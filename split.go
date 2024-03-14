package piscine

func Split(s, sep string) []string {
	var t []string
	r := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			t = append(t, s[r:i])
			r = i + len(sep)
			i += len(sep) - 1
		}
	}
	t = append(t, s[r:])
	return t
}
