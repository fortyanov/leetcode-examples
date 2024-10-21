package data_types

type BinarySearchTreeNode struct {
	Value int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func (bst *BinarySearchTree) Insert(val int) {
	if bst.Root == nil {
		bst.Root = &BinarySearchTreeNode{Value: val}
		return
	}

	bst.InsertRec(bst.Root, val)
}

func (bst *BinarySearchTree) InsertRec(node *BinarySearchTreeNode, val int) *BinarySearchTreeNode {
	if node == nil {
		return &BinarySearchTreeNode{Value: val}
	}

	if val <= node.Value {
		node.Left = bst.InsertRec(node.Left, val)
	}

	if val > node.Value {
		node.Right = bst.InsertRec(node.Right, val)
	}

	return node
}

func (bst *BinarySearchTree) Search(val int) bool {
	found := bst.SearchRec(bst.Root, val)

	return found
}

func (bst *BinarySearchTree) SearchRec(node *BinarySearchTreeNode, val int) bool {
	if node == nil {
		return false
	}

	if node.Value == val {
		return true
	}

	if val < node.Value {
		return bst.SearchRec(node.Left, val)
	}

	if val > node.Value {
		return bst.SearchRec(node.Right, val)
	}

	return false
}

func (bst *BinarySearchTree) GetValuesInorder(node *BinarySearchTreeNode, ch chan int) {
	if node == nil {
		return
	}

	bst.GetValuesInorder(node.Left, ch)

	//fmt.Print(node.Value, " ")
	ch <- node.Value

	bst.GetValuesInorder(node.Right, ch)
}
