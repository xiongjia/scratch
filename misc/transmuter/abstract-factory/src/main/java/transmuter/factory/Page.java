package transmuter.factory;

import java.util.ArrayList;
import java.util.List;

public abstract class Page {
  protected final String title;
  protected final String author;
  protected final List<Item> contents = new ArrayList<>();

  public Page(String title, String author) {
    this.title = title;
    this.author = author;
  }

  public void add(Item item) {
    contents.add(item);
  }

  public abstract String makeHtml();
}
