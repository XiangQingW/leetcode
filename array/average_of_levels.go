package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	r := []float64{}
	if root == nil {
		return r
	}

	cur_levels := []*TreeNode{root}
	next_levels := []*TreeNode{}

	for len(cur_levels) != 0 {
		sum := 0
		for _, node := range cur_levels {
			sum += node.Val

			if node.Left != nil {
				next_levels = append(next_levels, node.Left)
			}
			if node.Right != nil {
				next_levels = append(next_levels, node.Right)
			}
		}

		avg := float64(sum) / float64(len(cur_levels))
		r = append(r, avg)

		cur_levels = next_levels
		next_levels = []*TreeNode{}
	}

	return r
}
