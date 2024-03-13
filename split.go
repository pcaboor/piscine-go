package piscine

func Split(s, sep string) []string {
	var t []string
	st := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			t = append(t, s[st:i])
			st = i + len(sep)
			i += len(sep) - 1
		}
	}
	t = append(t, s[st:])
	return t
}
