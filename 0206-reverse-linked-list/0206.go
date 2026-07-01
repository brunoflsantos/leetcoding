package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Iterative solution: O(n)

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

// Recursive solution: O(n)

// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     newHead := reverseList(head.Next)
//     head.Next.Next = head
//     head.Next = nil
//     return newHead
// }
