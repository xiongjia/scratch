package transmuter.listfactory;

import transmuter.factory.Link;

public class ListLink extends Link {
  public ListLink(String caption, String url) {
    super(caption, url);
  }

  @Override
  public String makeHtml() {
    return String.format("  <li><a href=\"%s\">%s</a></li>\n", this.url, this.caption);
  }
}
