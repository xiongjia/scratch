package probius.lintcode;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertNull;

import org.junit.Test;

public class ListNodeTest {
  @Test
  public void test() {
    ListNode node = ListNode.create(null);
    assertNull(node);
    assertEquals(ListNode.dump(node), "<null>");

    node = ListNode.create(new int[] {});
    assertNull(node);
    assertEquals(ListNode.dump(node), "<null>");

    node = ListNode.create(new int[] { 1, 2 });
    assertNotNull(node);
    assertEquals(node.dump(), "([0]={1})->([1]={2})-><null>");
  }
}
