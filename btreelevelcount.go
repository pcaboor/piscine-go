package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lft := BTreeLevelCount(root.Left)
	rgt := BTreeLevelCount(root.Right)

	if lft > rgt {
		return lft + 1
	} else {
		return rgt + 1
	}
}
