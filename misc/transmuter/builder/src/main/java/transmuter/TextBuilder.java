package transmuter;

public class TextBuilder extends Builder {
  private final StringBuffer buffer = new StringBuffer();

  @Override
  public void makeTitle(String title) {
    buffer.append("Title: ").append(title).append("\n");
  }

  @Override
  public void makeString(String str) {
    buffer.append("> ").append(str).append("\n");
  }

  @Override
  public void makeItems(String[] items) {
    for (final String item : items) {
      buffer.append("- ").append(item).append("\n");
    }
  }

  @Override
  public void close() {
    buffer.append("~~~");
  }

  @Override
  public String getResult() {
    return buffer.toString();
  }
}
