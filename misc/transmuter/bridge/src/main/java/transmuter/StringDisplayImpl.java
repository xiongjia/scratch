package transmuter;

public class StringDisplayImpl extends DisplayImpl {
  private final String source;
  private final int sourceWidth;

  public StringDisplayImpl(String source) {
    this.source = source;
    this.sourceWidth = source.getBytes().length;
  }

  @Override
  public void rawOpen() {
    printLine();
  }

  @Override
  public void rawPrint() {
    System.out.println("|" + this.source + "|");
  }

  @Override
  public void rawClose() {
    printLine();
  }

  private void printLine() {
    System.out.print("+");
    for (int i = 0; i < this.sourceWidth; i++) {
      System.out.print("-");
    }
    System.out.println("+");
  }
}
