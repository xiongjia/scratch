package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** decorator tests. */
  public static void main(String[] args) {
    log.debug("Decorator tests");
    final Display stringDisplay = new StringDisplay("String display");
    final Display sideBorder = new SideBorder(stringDisplay, '#');
    final Display fullBorder = new FullBorder(sideBorder);
    sideBorder.show();
    fullBorder.show();
  }
}
