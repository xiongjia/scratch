package probius.lintcode;

import org.junit.Test;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class MergeTwoSortedListsTest {
  private static final Logger log = LoggerFactory.getLogger(MergeTwoSortedListsTest.class);
  private MergeTwoSortedLists inst = new MergeTwoSortedLists();

  @Test
  public void test() {
    final ListNode l1 = ListNode.create(new int[]{1, 3, 8, 11, 15});
    log.debug(l1.dump());
    final ListNode l2 = ListNode.create(new int[]{2});
    log.debug(l2.dump());
    final ListNode l3 = inst.mergeTwoLists(l1, l2);
    log.debug(l3.dump());
  }
}
