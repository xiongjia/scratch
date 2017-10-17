package probius.lintcode;

/**
 * The ListNode data class for lintcode exercise.
 */
public class ListNode {
  public int val;
  public ListNode next;

  public ListNode(int val) {
    this.val = val;
    this.next = null;
  }

  /** create list nodes. */
  public static ListNode create(int[] data) {
    if (data == null) {
      return null;
    }

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

  public String dump() {
    return ListNode.dump(this);
  }

  /** dump list nodes. */
  public static String dump(ListNode hdr) {
    if (hdr == null) {
      return "<null>";
    }
    final StringBuilder sb = new StringBuilder();
    int idx = 0;
    for (ListNode item = hdr; item != null; item = item.next) {
      sb.append(String.format("([%d]={%d})->", idx++, item.val));
    }
    sb.append("<null>");
    return sb.toString();
  }
}
