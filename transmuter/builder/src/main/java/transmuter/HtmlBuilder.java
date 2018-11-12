package transmuter;

public class HtmlBuilder extends Builder {
  private final StringBuffer buffer = new StringBuffer();

  @Override
  public void makeTitle(String title) {
    buffer.append("<h>").append(title).append("</h>");
  }

  @Override
  public void makeString(String str) {
    buffer.append("<p>").append(str).append("</p>");
  }

  @Override
  public void makeItems(String[] items) {
    buffer.append("<ul>");
    for (final String item : items) {
      buffer.append("<li>").append(item).append("</li>");
    }
    buffer.append("</ul>");
  }

  @Override
  public void close() {
    buffer.append("<br />");
  }

  @Override
  public String getResult() {
    return buffer.toString();
  }
}
