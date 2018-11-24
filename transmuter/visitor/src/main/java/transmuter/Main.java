package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Visitor tests. */
  public static void main(String[] args) {
    log.debug("Visitor tests");

    final Directory root = new Directory("root");
    root.add(new File("file1", 100));
    root.accept(new ListVisitor());
  }
}
