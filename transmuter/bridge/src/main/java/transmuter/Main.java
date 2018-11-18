package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** bridge tests. */
  public static void main(String[] args) {
    log.debug("Bridge tests");

    final Display display = new Display(new StringDisplayImpl("test1"));
    display.display();

    final CountDisplay countDisplay = new CountDisplay(new StringDisplayImpl("test2"));
    countDisplay.multiDisplay(3);
  }
}
