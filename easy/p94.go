package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func inorderTraversal(root *TreeNode) []int {
	var ans []int = make([]int,0,10)
	traversal(root,&ans)
	return ans
}

func traversal(node *TreeNode,ans *[]int){
	if node != nil{
		traversal(node.Left,ans)
		*ans = append(*ans, node.Val)
		traversal(node.Right,ans)
	}
}