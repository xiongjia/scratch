package probius.lintcode;

import org.junit.Test;

public class MergeTwoSortedListsTest {
  private MergeTwoSortedLists inst = new MergeTwoSortedLists();

  @Test
  public void test() {
    final ListNode l1 = ListNode.create(new int[]{1, 3, 8, 11, 15});
    l1.dump();
    final ListNode l2 = ListNode.create(new int[]{2});
    l2.dump();
    final ListNode l3 = inst.mergeTwoLists(l1, l2);
    l3.dump();
  }
}
