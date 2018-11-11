package transmuter.framework;

import java.util.HashMap;

public class Manager {
  private final HashMap<String, Product> showcase = new HashMap<>();

  public void register(String name, Product product) {
    showcase.put(name, product);
  }

  public Product create(String name) {
    final Product product = showcase.get(name);
    return product.createClone();
  }
}
