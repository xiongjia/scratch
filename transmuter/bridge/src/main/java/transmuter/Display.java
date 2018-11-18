package transmuter;

public class Display {
  private final DisplayImpl impl;

  public Display(DisplayImpl impl) {
    this.impl = impl;
  }

  public void open() {
    impl.rawOpen();
  }

  public void close() {
    impl.rawClose();
  }

  public void print() {
    impl.rawPrint();
  }

  /** Display. */
  public final void display() {
    open();
    print();
    close();
  }
}
