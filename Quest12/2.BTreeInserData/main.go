package main

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}
	
}
