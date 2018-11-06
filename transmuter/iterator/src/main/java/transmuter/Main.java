package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Iterator tests. */
  public static void main(String[] args) {
    log.debug("Iterator tests");

    final BookShelf bookShelf = new BookShelf();
    bookShelf.appendBook(new Book("Book 1"));
    bookShelf.appendBook(new Book("Book 2"));
    bookShelf.appendBook(new Book("Book 3"));
    final Iterator iterator = bookShelf.iterator();
    while (iterator.hasNext()) {
      final Book item = (Book)iterator.next();
      log.debug("Item {}", item.toString());
    }
  }
}
