package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data.(int) >= root.Data.(int) { 	// on check si la branche gauche vérifie bien les conditions de ABR que les enfants soient inférieures à la racine
		return false
	}
	if root.Right != nil && root.Right.Data.(int) <= root.Data.(int) { 	// on check si la branche gauche vérifie bien les conditions de ABR que les enfants soient supèrieures à la racine
		return false
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
