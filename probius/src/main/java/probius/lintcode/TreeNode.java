package probius.lintcode;

import java.util.StringTokenizer;

/**
 * The TreeNode data class for lintcode exercise.
 */
public class TreeNode {
  public int val;
  public TreeNode left = null;
  public TreeNode right = null;

  public TreeNode(int val) {
    this.val = val;
  }

  public String serialize() {
    return serialize(this);
  }

  /** serialize. */
  public static String serialize(TreeNode root) {
    if (root == null) {
      return "";
    }
    final StringBuilder sb = new StringBuilder();
    serialize(root, sb);
    return sb.toString();
  }

  private static void  serialize(TreeNode root, StringBuilder sb) {
    if (root == null) {
      sb.append("#,");
      return;
    }
    sb.append(String.format("%d,", root.val));
    serialize(root.left, sb);
    serialize(root.right, sb);
  }

  /** deserialize. */
  public static TreeNode deserialize(String tree) {
    if (tree == null || tree.isEmpty()) {
      return null;
    }
    return deserialize(new StringTokenizer(tree, ","));
  }

  private static TreeNode deserialize(StringTokenizer tokenizer) {
    if (!tokenizer.hasMoreTokens()) {
      return null;
    }
    final String node = tokenizer.nextToken().trim();
    if (node.equals("#")) {
      return null;
    }

    final TreeNode newNode = new TreeNode(Integer.parseInt(node));
    newNode.left = deserialize(tokenizer);
    newNode.right = deserialize(tokenizer);
    return newNode;
  }
}
