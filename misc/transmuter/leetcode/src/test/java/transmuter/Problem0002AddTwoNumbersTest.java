package transmuter;

import org.junit.Assert;
import org.junit.Test;
import transmuter.misc.ListNode;

public class Problem0002AddTwoNumbersTest {
  private Problem0002AddTwoNumbers problem = new Problem0002AddTwoNumbers();

  @Test
  public void addTwoNumbersTest1() {
    final ListNode l1 = ListNode.make(new int[] {2, 4, 3});
    final ListNode l2 = ListNode.make(new int[] {5, 6, 4});
    final ListNode result = problem.addTwoNumbers(l1, l2);
    Assert.assertTrue(ListNode.compare(result, new int[] {7, 0, 8}));
  }

  @Test
  public void addTwoNumbersTest2() {
    final ListNode l1 = ListNode.make(new int[] {0});
    final ListNode l2 = ListNode.make(new int[] {0});
    final ListNode result = problem.addTwoNumbers(l1, l2);
    Assert.assertTrue(ListNode.compare(result, new int[] {0}));
  }

  @Test
  public void addTwoNumbersTest3() {
    final ListNode l1 = ListNode.make(new int[] {1, 8});
    final ListNode l2 = ListNode.make(new int[] {0});
    final ListNode result = problem.addTwoNumbers(l1, l2);
    System.out.println(ListNode.makeDebugString(result));
    Assert.assertTrue(ListNode.compare(result, new int[] {1, 8}));
  }
}
