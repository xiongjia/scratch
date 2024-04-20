package transmuter;

import transmuter.framework.Product;

public class MessageBox extends Product {
  private final char boxchar;

  public MessageBox(char boxchar) {
    this.boxchar = boxchar;
  }

  @Override
  public void use(String str) {
    final int length = str.getBytes().length;
    for (int i = 0; i < length + 4; i++) {
      System.out.print(boxchar);
    }
    System.out.print("\n");
    System.out.println(boxchar + " "  + str + " " + boxchar);
    for (int i = 0; i < length + 4; i++) {
      System.out.print(boxchar);
    }
    System.out.print("\n");
  }
}
