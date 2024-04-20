package transmuter;

public class SideBorder extends Border {
  private final char border;

  public SideBorder(Display display, char border) {
    super(display);
    this.border = border;
  }

  @Override
  public int getColumns() {
    return display.getColumns() + 2;
  }

  @Override
  public int getRows() {
    return display.getRows();
  }

  @Override
  public String getRowText(int row) {
    return border + display.getRowText(row) + border;
  }
}
