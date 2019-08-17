package transmuter;

import transmuter.misc.ListNode;

/*
 * Problem 2: Add Two Numbers (medium) https://leetcode.com/problems/add-two-numbers/
 *
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 *
 * Example:
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 */
public class Problem0002AddTwoNumbers {
  /** Add two numbers. */
  public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
    ListNode hdr1 = l1;
    ListNode hdr2 = l2;
    ListNode hdr = null;
    ListNode tail = null;
    int carry = 0;
    while (!(hdr1 == null && hdr2 == null)) {
      final int val1 = hdr1 ==  null ? 0 : hdr1.val;
      final int val2 = hdr2 == null ? 0 : hdr2.val;
      hdr1 = hdr1 == null ? null : hdr1.next;
      hdr2 = hdr2 == null ? null : hdr2.next;

      final int num = val1 + val2 + carry;
      final int val = num % 10;
      carry = val != num ? 1 : 0;
      final ListNode node = new ListNode(val);
      if (hdr == null) {
        hdr = node;
        tail = node;
      } else {
        tail.next = node;
        tail = node;
      }
    }
    if (carry != 0) {
      final ListNode node = new ListNode(1);
      if (hdr == null) {
        hdr = node;
      } else {
        tail.next = node;
      }
    }

    return hdr == null ? new ListNode(0) : hdr;
  }

}
