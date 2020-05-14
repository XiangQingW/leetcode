package main

func getDistanceK(node *TreeNode, d int, K int, res *[]int) {
	if node == nil {
		return
	}

	if d == K {
		*res = append(*res, node.Val)
		return
	}

	d++
	getDistanceK(node.Left, d, K, res)
	getDistanceK(node.Right, d, K, res)
}

func getPath(node *TreeNode, target *TreeNode, path *[]*TreeNode) {
	if node == nil {
		return
	}

	*path = append(*path, node)
	if node == target {
		return
	}

	old_len := len(*path)
	getPath(node.Left, target, path)
	if (*path)[len(*path)-1] != target {
		*path = (*path)[:old_len]
	}

	old_len = len(*path)
	getPath(node.Right, target, path)
	if (*path)[len(*path)-1] != target {
		*path = (*path)[:old_len]
	}
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	res := []int{}
	if K == 0 {
		return append(res, target.Val)
	}

	getDistanceK(target, 0, K, &res)

	target.Left = nil
	target.Right = nil

	path := []*TreeNode{}
	getPath(root, target, &path)

	for i := len(path) - 1; i > 0; i-- {
		cur := path[i]
		pre := path[i-1]

		if cur.Left == nil {
			cur.Left = pre
		} else {
			cur.Right = pre
		}

		if pre.Left == cur {
			pre.Left = nil
		} else {
			pre.Right = nil
		}
	}

	getDistanceK(target, 0, K, &res)

	return res
}
