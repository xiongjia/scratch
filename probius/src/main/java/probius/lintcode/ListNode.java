package probius.lintcode;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * The ListNode data class for lintcode exercise.
 */
public class ListNode {
  private static final Logger log = LoggerFactory.getLogger(ListNode.class);

  public int val;
  public ListNode next;

  public ListNode(int val) {
    this.val = val;
    this.next = null;
  }

  /** create list nodes. */
  public static ListNode create(int[] data) {
    ListNode result = null;
    ListNode tail = null;
    for (int item : data) {
      ListNode node = new ListNode(item);
      if (result == null) {
        result = node;
        tail = node;
      } else {
        tail.next = node;
        tail = node;
      }
    }
    return result;
  }

  public void dump() {
    ListNode.dump(this);
  }

  /** dump list nodes. */
  public static void dump(ListNode hdr) {
    if (hdr == null) {
      log.debug("<null>");
    }
    final StringBuilder sb = new StringBuilder();
    int idx = 0;
    for (ListNode item = hdr; item != null; item = item.next) {
      sb.append(String.format("([%d]={%d})->", idx++, item.val));
    }
    sb.append("<null>");
    log.debug(sb.toString());
  }
}
