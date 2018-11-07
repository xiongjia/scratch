package transmuter;

public class BookShelfIterator implements Iterator {
  private final BookShelf bookShelf;
  private int index;

  public BookShelfIterator(BookShelf bookShelf) {
    this.bookShelf = bookShelf;
    this.index = 0;
  }

  @Override
  public boolean hasNext() {
    return this.index < this.bookShelf.getLength();
  }

  @Override
  public Object next() {
    return bookShelf.getBookAt(this.index++);
  }
}
