package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Left.Data >= root.Data) || (root.Right != nil && root.Right.Data <= root.Data) { // on check si la branche gauche et droite vérifient bien les conditions de ABR que les enfants soient respectivement inférieurs et superieurs  à la racine
		return false
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
