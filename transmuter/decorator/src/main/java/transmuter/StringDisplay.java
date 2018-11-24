package transmuter;

public class StringDisplay extends Display {
  private final String source;

  public StringDisplay(String source) {
    this.source = source;
  }

  @Override
  public int getColumns() {
    return source.getBytes().length;
  }

  @Override
  public int getRows() {
    return 1;
  }

  @Override
  public String getRowText(int row) {
    return row == 0 ? source : null;
  }
}
