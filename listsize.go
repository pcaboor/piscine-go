package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
count := 0
  current := l.Head

 for current != nil {
		current = current.Next
   count++
	}
  return count
}
