package transmuter;

public class CountDisplay extends Display {
  public CountDisplay(DisplayImpl impl) {
    super(impl);
  }

  /** Multi-Display. */
  public void multiDisplay(int times) {
    open();
    for (int i = 0; i < times; i++) {
      print();
    }
    close();
  }
}
