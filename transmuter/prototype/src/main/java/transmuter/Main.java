package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import transmuter.framework.Manager;
import transmuter.framework.Product;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Prototype tests. */
  public static void main(String[] args) {
    log.debug("Prototype tests");

    final Manager manager = new Manager();
    manager.register("strong", new UnderlinePen('~'));
    manager.register("slash", new MessageBox('/'));
    manager.register("warning", new MessageBox('*'));

    final Product productStrong = manager.create("strong");
    final Product productSlash = manager.create("slash");
    final Product productWarn = manager.create("warning");
    productStrong.use("Strong");
    productSlash.use("slash");
    productWarn.use("warning");
  }
}
