package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import transmuter.framework.Factory;
import transmuter.framework.Product;
import transmuter.idcard.IdCardFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Factory method tests. */
  public static void main(String[] args) {
    log.debug("Factory method tests");
    final Factory factory = new IdCardFactory();
    final Product card1 = factory.create("owner1");
    final Product card2 = factory.create("owner2");
    card1.use();
    card2.use();
  }
}
