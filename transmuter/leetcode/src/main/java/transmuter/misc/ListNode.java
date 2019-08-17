package transmuter.misc;

import java.util.List;

public class ListNode {
  public int val;
  public ListNode next;

  public ListNode(int val) {
    this.val = val;
  }

  /** Make list from array. */
  public static ListNode make(int[] src) {
    if (src == null) {
      return null;
    }
    ListNode hdr = null;
    ListNode tail = null;
    for (final int val : src) {
      final ListNode node = new ListNode(val);
      if (hdr == null) {
        hdr = node;
        tail = node;
      } else {
        tail.next = node;
        tail = node;
      }
    }
    return hdr;
  }

  /** Make debug string. */
  public static String makeDebugString(ListNode hdr) {
    final StringBuilder sb = new StringBuilder();
    for (ListNode node = hdr; node != null; node = node.next) {
      sb.append(node.val).append(node.next == null ? "" : " -> ");
    }
    return sb.toString();
  }

  /** Compare ListNode vs Arrary. */
  public static boolean compare(ListNode src1, int[] src2) {
    final int src2Size = src2 == null ? 0 : src2.length;
    int index = 0;
    for (ListNode node = src1; node != null; node = node.next, index++) {
      if (index >= src2Size) {
        return false;
      }
      if (node.val != src2[index]) {
        return false;
      }
    }
    return index == src2Size;
  }
}
