package transmuter;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class Directory extends Entry {
  private final String name;
  private final List<Entry> children = new ArrayList<>();

  public Directory(String name) {
    this.name = name;
  }

  @Override
  public String getName() {
    return this.name;
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
  protected void printList(String prefix) {
    System.out.println(prefix + "/" + this);
    final Iterator itr = children.iterator();
    final String directoryPrefix = prefix + "/" + name;
    while (itr.hasNext()) {
      final Entry entry = (Entry)itr.next();
      entry.printList(directoryPrefix);
    }
  }

  public Entry add(Entry entry) {
    children.add(entry);
    return this;
  }
}
