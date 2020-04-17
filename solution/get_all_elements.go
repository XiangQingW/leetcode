package main

// import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeft(root *TreeNode, s *[]*TreeNode) {
	for root != nil {
		*s = append(*s, root)
		root = root.Left
	}
}

func traverse(s *[]*TreeNode) {
	last_node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	node := last_node.Right
	findLeft(node, s)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	r, s1, s2 := []int{}, []*TreeNode{}, []*TreeNode{}
	findLeft(root1, &s1)
	findLeft(root2, &s2)

	for len(s1) != 0 || len(s2) != 0 {
		isNode1 := false
		var node1, node2 *TreeNode
		if 0 < len(s1) {
			node1 = s1[len(s1)-1]
		}
		if 0 < len(s2) {
			node2 = s2[len(s2)-1]
		}
		if node1 == nil {
			isNode1 = false
		} else if node2 == nil {
			isNode1 = true
		} else if node1.Val < node2.Val {
			isNode1 = true
		} else {
			isNode1 = false
		}

		if isNode1 {
			r = append(r, node1.Val)
			traverse(&s1)
		} else {
			r = append(r, node2.Val)
			traverse(&s2)
		}

	}

	return r
}

func main() {

}
