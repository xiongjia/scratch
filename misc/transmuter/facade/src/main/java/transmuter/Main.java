package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import transmuter.pagemaker.PageMaker;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Facade tests. */
  public static void main(String[] args) {
    log.debug("Facade tests");
    log.debug("Page: {}", PageMaker.makeWelcomePage("test1@test.com"));
  }
}
