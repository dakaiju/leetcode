// https://leetcode.com/problems/diameter-of-binary-tree/

package solutions

import "kuang/leetcode/data"

type TreeNode = data.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	sub, _ := work(root)
	return sub
}

// return: (max in subtree, max from leaf)
func work(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lm, ll := work(root.Left)
	rm, rl := work(root.Right)

	return max(max(ll+rl, lm), rm), max(ll, rl) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
