/* You are given two non-empty linked lists representing two non-negative integers.
   The digits are stored in reverse order, and each of their nodes contains a single digit.
   Add the two numbers and return the sum as a linked list.
   You may assume the two numbers do not contain any leading zero, except the number 0 itself.

   Example 1:
     Input: l1 = [2,4,3], l2 = [5,6,4]
	 Output: [7,0,8]
	 Explanation: 342 + 465 = 807.


   Definition for singly-linked list.
   type ListNode struct {
       Val int
       Next *ListNode
   }
*/

/* Вторая попытка решить задачу, путем цикла вместо рекурсии.
   Если я прав, то рекурсия кладет все в стек и это замедляет решение на больших массивах данных.

   Вышло 6 мс по скорости.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	carry := 0

	for l1 != nil || l2 != nil {
		a, b := 0, 0
		temp := 0

		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		temp = a + b + carry
		current.Next = &ListNode{Val: temp % 10}
		current = current.Next
		carry = temp / 10
	}

	// На последней итерации l1 и l2 становятся nil, поэтому нужно проверить есть ли избыток в carry и дополнить список
	if carry == 1 {
		current.Next = &ListNode{Val: 1}
	}

	return result.Next
}