package transmuter;

public class CharDisplay extends AbstractDisplay {
  private final char source;

  public CharDisplay(char source) {
    this.source = source;
  }

  @Override
  protected void open() {
    System.out.print("<<");
  }

  @Override
  protected void close() {
    System.out.println(">>");
  }

  @Override
  protected void print() {
    System.out.print(this.source);
  }
}
