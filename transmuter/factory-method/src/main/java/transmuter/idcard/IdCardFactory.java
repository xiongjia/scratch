package transmuter.idcard;

import java.util.HashMap;

import transmuter.framework.Factory;
import transmuter.framework.Product;

public class IdCardFactory extends Factory {
  private final HashMap products = new HashMap();
  private int serialIndex = 10;

  @Override
  protected synchronized Product createProduct(String owner) {
    return new IdCard(owner, serialIndex++);
  }

  @Override
  protected void registerProduct(Product product) {
    final IdCard card = (IdCard)product;
    products.put(new Integer(card.getSerial()), card.getOwner());
  }
}
