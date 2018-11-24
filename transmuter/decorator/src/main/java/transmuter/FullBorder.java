package transmuter;

public class FullBorder extends Border {
  public FullBorder(Display display) {
    super(display);
  }

  @Override
  public int getColumns() {
    return display.getColumns() + 2;
  }

  @Override
  public int getRows() {
    return display.getRows() + 2;
  }

  @Override
  public String getRowText(int row) {
    return row == 0 || row == display.getRows() + 1
      ? ("+" + makeLine('-', display.getColumns()) + "+")
      : ("|" + display.getRowText(row - 1) + "|");
  }

  private String makeLine(char ch, int count) {
    final StringBuffer sb = new StringBuffer();
    for (int i = 0; i < count; i++) {
      sb.append(ch);
    }
    return sb.toString();
  }
}
