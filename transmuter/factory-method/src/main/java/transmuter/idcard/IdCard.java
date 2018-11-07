package transmuter.idcard;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import transmuter.framework.Product;

public class IdCard extends Product {
  private static final Logger log = LoggerFactory.getLogger(IdCard.class);

  private final String owner;
  private final int serial;

  IdCard(String owner, int serial) {
    this.owner = owner;
    this.serial = serial;
  }

  public String getOwner() {
    return owner;
  }

  public int getSerial() {
    return serial;
  }

  @Override
  public void use() {
    log.debug("using id card[{} - {}]", this.owner, this.serial);
  }
}
