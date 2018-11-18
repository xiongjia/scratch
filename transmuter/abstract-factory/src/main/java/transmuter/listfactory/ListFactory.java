package transmuter.listfactory;

import transmuter.factory.Factory;
import transmuter.factory.Link;
import transmuter.factory.Page;

public class ListFactory extends Factory {
  @Override
  public Link createLink(String caption, String url) {
    return new ListLink(caption, url);
  }

  @Override
  public Page createPage(String title, String author) {
    return new ListPage(title, author);
  }
}
