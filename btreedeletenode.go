package piscine 

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
      s := FindN(node.Right)
      node.Data = s.Data
      node.Right = BTreeDeleteNode(node.Right, s)
    }
  }
  return root
}

func FindN(node *TreeNode) *TreeNode {
  for node.Left != nil {
    node = node.Left
  }
  return node
}
