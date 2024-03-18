package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	current := l
	for i := 1; i <= pos; i++ {
		if current.Next == nil {
			return nil
		}
		current = current.Next
	}
	return current
}
