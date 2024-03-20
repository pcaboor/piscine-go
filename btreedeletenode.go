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
