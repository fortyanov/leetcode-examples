package main

import "leetcode-examples/utils/data_types"

/*
inp: 2->3->4

res: 4->3->2
*/

func reverseLinkedList(inp *data_types.LinkedList) {
	var prev *data_types.LinkedListNode

	curr := inp.Head

	for curr != nil {
		next := curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	inp.Head = prev
}

func main() {
	inp := new(data_types.LinkedList)
	inp.Insert(2)
	inp.Insert(3)
	inp.Insert(4)

	reverseLinkedList(inp)
	inp.PrintValues(inp.Head)
}
