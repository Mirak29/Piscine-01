package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rightCpt := BTreeLevelCount(root.Right)
	leftCpt := BTreeLevelCount(root.Left)

	if leftCpt > rightCpt {
		return leftCpt + 1
	}
	return rightCpt + 1
}
