  package main


 // ListNode Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{}
	current := result
	node1 := l1
	node2 := l2
	carry := 0

	for node1 != nil || node2 != nil {
		x := 0
		if node1 != nil {
			x = node1.Val
			node1 = node1.Next
		}

		y := 0
		if node2 != nil {
			y = node2.Val
			node2 = node2.Next
		}

		sum := carry + x + y
		val := sum % 10
		carry = sum / 10

		resultNode := &ListNode{Val: val, Next: nil}
		current.Next = resultNode
		current = current.Next
	}

	if carry > 0 {
		resultNode := &ListNode{Val: carry, Next: nil}
		current.Next = resultNode
		current = current.Next
	}

	return result.Next

}
