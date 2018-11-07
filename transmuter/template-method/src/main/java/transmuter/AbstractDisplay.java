package transmuter;

public abstract class AbstractDisplay {
  protected abstract void open();

  protected abstract void close();

  protected abstract void print();

  /** Display. */
  public final void display() {
    open();
    for (int i = 0; i < 5; ++i) {
      print();
    }
    close();
  }
}
