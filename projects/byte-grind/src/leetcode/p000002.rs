// https://leetcode.com/problems/add-two-numbers

/*
* 2. Add Two Numbers
*
* [medium]
*
* You are given two non-empty linked lists representing two non-negative integers.
* The digits are stored in reverse order, and each of their nodes contains a single digit.
* Add the two numbers and return the sum as a linked list.
*
* You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*
* Example 1:
* Input: l1 = [2,4,3], l2 = [5,6,4]
* Output: [7,0,8]
* Explanation: 342 + 465 = 807.
*
* Example 2:
* Input: l1 = [0], l2 = [0]
* Output: [0]
*
* Example 3:
* Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
* Output: [8,9,9,9,0,0,0,1]
*
* Constraints:
* - The number of nodes in each linked list is in the range [1, 100].
* - 0 <= Node.val <= 9
* - It is guaranteed that the list represents a number that does not have leading zeros.
*/

use crate::util::linked_list::{to_list, ListNode};

pub struct Solution {}

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let (mut l1, mut l2) = (l1, l2);
        let mut dummy_head = Some(Box::new(ListNode::new(0)));
        let mut tail = &mut dummy_head;
        let (mut l1_end, mut l2_end, mut overflow) = (false, false, false);
        loop {
            let lhs = match l1 {
                Some(node) => {
                    l1 = node.next;
                    node.val
                }
                None => {
                    l1_end = true;
                    0
                }
            };
            let rhs = match l2 {
                Some(node) => {
                    l2 = node.next;
                    node.val
                }
                None => {
                    l2_end = true;
                    0
                }
            };
            // if l1, l2 end and there is not overflow from previous operation, return the result
            if l1_end && l2_end && !overflow {
                break dummy_head.unwrap().next;
            }
            let sum = lhs + rhs + if overflow { 1 } else { 0 };
            let sum = if sum >= 10 {
                overflow = true;
                sum - 10
            } else {
                overflow = false;
                sum
            };
            tail.as_mut().unwrap().next = Some(Box::new(ListNode::new(sum)));
            tail = &mut tail.as_mut().unwrap().next
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p2() {
        assert_eq!(
            Solution::add_two_numbers(to_list(vec![2, 4, 3]), to_list(vec![5, 6, 4])),
            to_list(vec![7, 0, 8])
        );

        assert_eq!(
            Solution::add_two_numbers(to_list(vec![9, 9, 9, 9]), to_list(vec![9, 9, 9, 9, 9, 9])),
            to_list(vec![8, 9, 9, 9, 0, 0, 1])
        );

        assert_eq!(
            Solution::add_two_numbers(to_list(vec![0]), to_list(vec![0])),
            to_list(vec![0])
        )
    }
}
