package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Template method tests. */
  public static void main(String[] args) {
    log.debug("Template method tests");

    final AbstractDisplay str = new StringDisplay("Test Strings");
    str.display();

    final AbstractDisplay ch = new CharDisplay('x');
    ch.display();
  }
}
