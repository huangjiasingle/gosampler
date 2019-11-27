package main

import "fmt"

func main() {
	node1 := SetTreeNode(10)
	node2 := SetTreeNode(9)
	node3 := SetTreeNode(20)
	node4 := SetTreeNode(15)
	node5 := SetTreeNode(35)

	node1.SetLeftNode(node2)
	node1.SetRightNode(node3)
	node3.SetLeftNode(node4)
	node3.SetRightNode(node5)

	PreTraversalBTres(node1)
	MidTraversalBTres(node1)
	NextTraversalBTres(node1)

	i := []int{2, 3, 1, 4, 5}
	root := &TreeRoot{}
	for _, v := range i {
		CreateBTree(root, v)
	}

	PreTraversalBTres(root.GetTreeRoot())
	MidTraversalBTres(root.GetTreeRoot())
	NextTraversalBTres(root.GetTreeRoot())

}

// TreeNode 树叶子节点
type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

// SetTreeNode set tree node val
func SetTreeNode(val interface{}) *TreeNode {
	return &TreeNode{Value: val}
}

// SetLeftNode 设置左叶子节点
func (t *TreeNode) SetLeftNode(n *TreeNode) {
	t.Left = n
}

// SetRightNode 设置右叶子节点
func (t *TreeNode) SetRightNode(n *TreeNode) {
	t.Right = n
}

// GetLeftNode 获取左叶子节点
func (t *TreeNode) GetLeftNode() *TreeNode {
	return t.Left
}

// GetRightNode 获取右叶子节点
func (t *TreeNode) GetRightNode() *TreeNode {
	return t.Right
}

// TreeRoot 跟节点 代表一个树
type TreeRoot struct {
	Root *TreeNode
}

// SetTreeRoot 设置跟节点
func (r *TreeRoot) SetTreeRoot(t *TreeNode) {
	r.Root = t
}

// GetTreeRoot 获取跟节点
func (r *TreeRoot) GetTreeRoot() *TreeNode {
	return r.Root
}

// PreTraversalBTres 先序遍历二叉树树
func PreTraversalBTres(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Value)
		PreTraversalBTres(t.GetLeftNode())
		PreTraversalBTres(t.GetRightNode())
	}
}

// MidTraversalBTres 中序遍历二叉树树
func MidTraversalBTres(t *TreeNode) {
	if t != nil {
		PreTraversalBTres(t.GetLeftNode())
		fmt.Println(t.Value)
		PreTraversalBTres(t.GetRightNode())
	}
}

// NextTraversalBTres 后序遍历二叉树树
func NextTraversalBTres(t *TreeNode) {
	if t != nil {
		PreTraversalBTres(t.GetLeftNode())
		PreTraversalBTres(t.GetRightNode())
		fmt.Println(t.Value)
	}
}

// 先序(根->左->右)，中序(左->根->右)，后序(左->右->根)。如果访问有孩子的节点，先处理孩子的，随后返回

// 如果比当前根节点要小，那么放到当前根节点左边，如果比当前根节点要大，那么放到当前根节点右边。

// CreateBTree 动态创建二叉树
func CreateBTree(root *TreeRoot, in interface{}) {
	if root.GetTreeRoot() == nil {
		root.SetTreeRoot(&TreeNode{Value: in})
	} else {

		tempNode := root.GetTreeRoot()
	B:
		if tempNode != nil {
			v := tempNode.Value.(int)
			i := in.(int)
			if i > v {
				if tempNode.GetRightNode() == nil {
					tempNode.SetRightNode(&TreeNode{Value: in})
				} else {
					tempNode = tempNode.GetRightNode()
					goto B
				}
			} else {
				if tempNode.GetLeftNode() == nil {
					tempNode.SetLeftNode(&TreeNode{Value: in})
				} else {
					tempNode = tempNode.GetLeftNode()
					goto B
				}
			}
		}
	}
}
