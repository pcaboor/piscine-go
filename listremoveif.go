package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for current := l.Head; current != nil; current = current.Next {
		if current.Data == data_ref {
			return
		}
	}
	return
}
