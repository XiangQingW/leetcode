package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	r := [][]int{}
	if root == nil {
		return r
	}

	preLevel := []*TreeNode{root}
	curLevel := []*TreeNode{}

	for 0 < len(preLevel) {
		v := []int{}
		for _, node := range preLevel {
			v = append(v, node.Val)
			if node.Left != nil {
				curLevel = append(curLevel, node.Left)
			}
			if node.Right != nil {
				curLevel = append(curLevel, node.Right)
			}
		}
		r = append(r, v)

		preLevel = curLevel
		curLevel = []*TreeNode{}

	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func main() {

}
