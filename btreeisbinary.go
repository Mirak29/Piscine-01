package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	} else if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	} else if root.Right != nil && root.Right.Data < root.Data {
		return false
	} else if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	return true
}
