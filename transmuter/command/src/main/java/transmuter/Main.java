package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Command tests. */
  public static void main(String[] args) {
    log.debug("Command tests");

    final Stock stock = new Stock("test", 10);
    final Broker broker = new Broker();
    broker.takeOrder(new BuyStock(stock));
    broker.takeOrder(new SellStock(stock));
    broker.placeOrders();
  }
}
