package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	} else if elem < root.Data {
		BTreeSearchItem(root.Left, elem)
	} else {
		BTreeSearchItem(root.Right, elem)
	}
	return root
}