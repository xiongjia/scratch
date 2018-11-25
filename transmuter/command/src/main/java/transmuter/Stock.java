package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Stock {
  private static final Logger log = LoggerFactory.getLogger(Stock.class);
  private final String name;
  private final int quantity;

  public Stock(String name, int quantity) {
    this.name = name;
    this.quantity = quantity;
  }

  public void buy() {
    log.debug("buy: {} [{}]", this.name, this.quantity);
  }

  public void sell() {
    log.debug("sell: {} [{}]", this.name, this.quantity);
  }
}
