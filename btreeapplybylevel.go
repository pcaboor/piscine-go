package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	t := []*TreeNode{root}
	for len(t) > 0 {

		n := t[0]
		t = t[1:]

		_, err := f(n.Data)
		
		if err != nil {
			return
		}
		if n.Left != nil {
			t = append(t, n.Left)
		}
		if n.Right != nil {
			t = append(t, n.Right)
		}
	}
}
