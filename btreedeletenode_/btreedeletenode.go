package piscine

type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {

		if node.Left == nil && node.Right == nil {
			root = nil
		} else if node.Left == nil {
			root = root.Right
		} else if node.Right == nil {
			root = root.Left
		} else {
			successor := FindMinNode(node.Right)
			node.Data = successor.Data
			node.Right = BTreeDeleteNode(node.Right, successor)
		}
	}

	return root
}

func FindMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || elem == root.Data {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Left.Data >= root.Data) || (root.Right != nil && root.Right.Data <= root.Data) { // on check si la branche gauche et droite vérifient bien les conditions de ABR que les enfants soient respectivement inférieurs et superieurs  à la racine
		return false
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
