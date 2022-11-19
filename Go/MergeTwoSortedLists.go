package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := NewNode(1, nil)
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)

	list2 := NewNode(2, nil)
	list2.Add(2)
	list2.Add(3)
	list2.Add(3)

	mergedList := mergeTwoLists(list1, list2)
	fmt.Println(mergedList.Traverse())
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := NewNode(0, nil)

	curNode1 := list1
	curNode2 := list2

	for (curNode1 != nil) && (curNode2 != nil) {
		if curNode1.Val <= curNode2.Val {
			result.Add(curNode1.Val)
			// result = append(result,curNode1.Val)
			curNode1 = curNode1.Next
		} else {
			result.Add(curNode2.Val)
			// result = append(result,curNode2.Val)
			curNode2 = curNode2.Next
		}
	}

	for curNode1 != nil {
		result.Add(curNode1.Val)
		// result = append(result,curNode1.Val)
		curNode1 = curNode1.Next
	}

	for curNode2 != nil {
		result.Add(curNode2.Val)
		// result = append(result,curNode2.Val)
		curNode2 = curNode2.Next
	}
	return result.Next
}

func NewNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func GetLastNode(node *ListNode) *ListNode {
	cur := node
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func (l *ListNode) Add(val int) {
	cur := GetLastNode(l)
	newNode := NewNode(val, nil)
	cur.Next = newNode
}

func (l *ListNode) String() string {
	return fmt.Sprint(l.Val)
}

func (l *ListNode) Traverse() []int {
	cur := l
	var result []int
	for cur.Next != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}
