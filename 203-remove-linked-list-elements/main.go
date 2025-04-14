package main

/*
- Algo:
- 
*/

 // Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
 func removeElements(head *ListNode, val int) *ListNode {
	// Handle the case where the head itself is the target
   for head != nil && head.Val == val {
	   head = head.Next
   }

   curr := head
   // Iterate through the rest of the list
   for curr != nil && curr.Next != nil {
	   if curr.Next.Val == val {
		   // Skip the node with the value 'val'
		   curr.Next = curr.Next.Next
	   } else {
		   curr = curr.Next
	   }
   }

   return head
}