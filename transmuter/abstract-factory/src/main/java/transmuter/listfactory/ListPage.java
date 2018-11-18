package transmuter.listfactory;

import transmuter.factory.Item;
import transmuter.factory.Page;

import java.util.Iterator;

public class ListPage extends Page {
  public ListPage(String title, String author) {
    super(title, author);
  }

  @Override
  public String makeHtml() {
    final StringBuffer buffer = new StringBuffer();
    buffer.append("<html><head><title>" + this.title + "</title></head>\n");
    buffer.append("<body>\n");
    buffer.append("<h1>" + this.author + "</h1>\n");
    buffer.append("<ul>\n");
    final Iterator it = this.contents.iterator();
    while (it.hasNext()) {
      final Item item = (Item)it.next();
      buffer.append(item.makeHtml());
    }
    buffer.append("</ul>\n");
    buffer.append("<hr><address>" + this.author + "</address>");
    buffer.append("</body></html>\n");
    return buffer.toString();
  }
}
