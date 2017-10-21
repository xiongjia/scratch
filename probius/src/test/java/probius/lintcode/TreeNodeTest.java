package probius.lintcode;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;

import org.junit.Test;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class TreeNodeTest {
  private static final Logger log = LoggerFactory.getLogger(TreeNodeTest.class);

  @Test
  public void test() {
    final String simpleTree = "1,2,#,#,3";
    final TreeNode root = TreeNode.deserialize(simpleTree);
    assertNotNull(root);
    assertEquals(root.val, 1);
    assertNotNull(root.left);
    assertEquals(root.left.val, 2);
    assertNotNull(root.right);
    assertEquals(root.right.val, 3);
    log.debug("root = {}, left = {}, right = {}",
        root.val, root.left.val, root.right.val);
    log.debug("data = {}", root.serialize());
  }
}
