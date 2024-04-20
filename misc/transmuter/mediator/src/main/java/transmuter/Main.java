package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Mediator tests. */
  public static void main(String[] args) {
    log.debug("Mediator tests");
    final User user1 = new User("user1");
    final User user2 = new User("user2");

    user1.sendMessage("test1");
    user2.sendMessage("test2");
  }
}
