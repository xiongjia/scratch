package transmuter;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class Directory extends Entry {
  private final String name;
  private final List<Entry> children = new ArrayList<Entry>();

  public Directory(String name) {
    this.name = name;
  }

  @Override
  public String getName() {
    return name;
  }

  @Override
  public int getSize() {
    int size = 0;
    final Iterator itr = children.iterator();
    while (itr.hasNext()) {
      final Entry entry = (Entry)itr.next();
      size += entry.getSize();
    }
    return size;
  }

  @Override
  public Entry add(Entry entry) {
    children.add(entry);
    return this;
  }

  @Override
  public Iterator iterator() {
    return children.iterator();
  }

  @Override
  public void accept(Visitor visitor) {
    visitor.visit(this);
  }
}
