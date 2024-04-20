package transmuter;

public class StringDisplay extends AbstractDisplay {
  private final String source;
  private final int width;

  public StringDisplay(String source) {
    this.source = source;
    this.width = source.getBytes().length;
  }

  @Override
  protected void open() {
    printLine();
  }

  @Override
  protected void close() {
    printLine();
  }

  @Override
  protected void print() {
    System.out.println("|" + this.source + "|");
  }

  private void printLine() {
    System.out.print("+");
    for (int i = 0; i < this.width; ++i) {
      System.out.print("-");
    }
    System.out.println("+");
  }
}
