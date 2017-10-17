package probius.lintcode;

/**
 * LintCode exercise - Merge Two Sorted Lists
 *  http://www.lintcode.com/en/problem/merge-two-sorted-lists/
 *  Merge two sorted (ascending) linked lists and return it as a new sorted list.
 *  The new sorted list should be made by splicing together the nodes of
 *  the two lists and sorted in ascending order.
 * Example
 *  Given 1->3->8->11->15->null, 2->null , return 1->2->3->8->11->15->null.
 */
public class MergeTwoSortedLists {
  /** solution. */
  public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
    if (l1 == null) {
      return l2;
    }
    if (l2 == null) {
      return l1;
    }

    ListNode result = null;
    ListNode tail = null;
    ListNode hdr1 = l1;
    ListNode hdr2 = l2;
    for (; !(hdr1 == null || hdr2 == null);) {
      ListNode node = null;
      if (hdr1.val > hdr2.val) {
        node = hdr2;
        hdr2 = hdr2.next;
      } else {
        node = hdr1;
        hdr1 = hdr1.next;
      }
      node.next = null;

      if (tail == null) {
        result = node;
        tail = node;
      } else {
        tail.next = node;
        tail = node;
      }
    }
    if (hdr1 != null || hdr2 != null) {
      tail.next = (hdr1 == null ? hdr2 : hdr1);
    }
    return result;
  }
}
