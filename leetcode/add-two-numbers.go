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

// Комментарий от меня: решение оказалось неэффективным, я хотел не решать в лоб циклом, а попробовать иной способ, но вышло медленнее на 4 мс.
// UPD 12.02.2024 17:19 - Я объявил переменные в начале функции, вместо их объявления в каждом блоке if, ИМХО это улучшило логику, а еще ускорило код на 1 мс.

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	a, b := 0, 0
	temp := 0
	node := &ListNode

	if l1 != nil && l2 != nil {
		a, n1 = l1.Val, l1.Next
		b, n2 = l2.Val, l2.Next

		temp = a + b + carry

		node.Val = temp % 10
		node.Next = add(n1, n2, temp/10)
		return node
	} else if l1 != nil && l2 == nil {
		a, n1 = l1.Val, l1.Next

		temp = a + carry

		node.Val = temp % 10
		node.Next = add(n1, nil, temp/10)
		return node
	} else if l1 == nil && l2 != nil {
		b, n2 = l2.Val, l2.Next

		temp = b + carry

		node.Val = temp % 10
		node.Next = add(nil, n2, temp/10)
		return node
	} else {
		if carry != 0 {
			return &ListNode{Val: carry}
		} else {
			return nil
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}
