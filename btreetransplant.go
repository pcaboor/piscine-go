package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == node {
		return rplc
	}
	p := node.Parent
	if node == p.Left {
		p.Left = rplc
	} else {
		p.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = p
	}
	return root
}
