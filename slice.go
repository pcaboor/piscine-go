package piscine

func Slice(a []string, nbrs ...int) []string {
	var start, end int

	switch len(nbrs) {
	case 0:
		return nil
	case 1:
		start = nbrs[0]
		end = len(a)
	case 2:
		start = nbrs[0]
		end = nbrs[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}

	if start < 0 {
		start = 0
	}
	if end > len(a) {
		end = len(a)
	}

	if start >= end || start >= len(a) {
		return nil
	}
	return a[start:end]
}
