package main

import (
	"leetcode-examples/utils/data_types"
	"sync"
)

/*
https://leetcode.com/problems/merge-k-sorted-lists/description/

23. Merge k Sorted Lists
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:

Input: lists = []
Output: []

Example 3:

Input: lists = [[]]
Output: []

*/

func mergeKLists(lists []*data_types.LinkedList) (res *data_types.LinkedList) {
	res = new(data_types.LinkedList)
	bst := data_types.BinarySearchTree{}
	llValuesChan := make(chan int)
	bstValuesChan := make(chan int)
	wg := new(sync.WaitGroup)

	go func() {
		wg.Add(1)
		defer wg.Done()

		for _, l := range lists {
			l.GetValues(l.Head, llValuesChan)
		}
		close(llValuesChan)
	}()

	for v := range llValuesChan {
		bst.Insert(v)
	}

	go func() {
		wg.Add(1)
		defer wg.Done()

		bst.GetValuesInorder(bst.Root, bstValuesChan)
		close(bstValuesChan)
	}()

	for v := range bstValuesChan {
		res.Insert(v)
	}

	wg.Wait()

	return res
}

func main() {
	ll1 := new(data_types.LinkedList)
	ll1.Insert(1)
	ll1.Insert(4)
	ll1.Insert(5)

	ll2 := new(data_types.LinkedList)
	ll2.Insert(1)
	ll2.Insert(3)
	ll2.Insert(4)

	ll3 := new(data_types.LinkedList)
	ll3.Insert(2)
	ll3.Insert(6)

	llMerged := mergeKLists([]*data_types.LinkedList{ll1, ll2, ll3})
	llMerged.PrintValues(llMerged.Head)
}
