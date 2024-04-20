package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Observer tests. */
  public static void main(String[] args) {
    log.debug("Observer tests");

    final NumberGenerator generator = new RandomNumberGenerator();
    generator.addObserver(new DigitObserver());
    generator.execute();
  }
}
