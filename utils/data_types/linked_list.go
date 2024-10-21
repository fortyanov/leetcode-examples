package data_types

import "fmt"

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Root *LinkedListNode
}

func (ll *LinkedList) Insert(val int) {
	if ll.Root == nil {
		ll.Root = &LinkedListNode{Value: val}
		return
	}

	ll.InsertRec(ll.Root, val)
}

func (ll LinkedList) InsertRec(node *LinkedListNode, val int) *LinkedListNode {
	if node == nil {
		return &LinkedListNode{Value: val}
	}

	node.Next = ll.InsertRec(node.Next, val)

	return node
}

func (ll *LinkedList) GetValues(node *LinkedListNode, ch chan int) {
	if node == nil {
		return
	}

	ch <- node.Value

	ll.GetValues(node.Next, ch)
}

func (ll *LinkedList) PrintValues(node *LinkedListNode) {
	if node == nil {
		return
	}

	fmt.Print(node.Value, " ")

	ll.PrintValues(node.Next)
}
