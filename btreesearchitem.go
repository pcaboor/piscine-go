package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || elem == root.Data {
		return root
	} 
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}