package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Composite tests. */
  public static void main(String[] args) {
    log.debug("Composite tests");

    final Directory root = new Directory("root");
    final Directory dir1 = new Directory("dir1");
    final File file1 = new File("file1", 100);
    root.add(dir1.add(file1));
    log.debug("Root = {} [{}]", root.getName(), root.getSize());
    root.printList();
  }
}
