package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Interpreter tests. */
  public static void main(String[] args) {
    log.debug("Interpreter tests");

    final Expression isMale = new OrExpression(new TerminalExpression("Male1"),
        new TerminalExpression("Male2"));
    log.debug("Test: {}", isMale.interpret("test"));
  }
}
