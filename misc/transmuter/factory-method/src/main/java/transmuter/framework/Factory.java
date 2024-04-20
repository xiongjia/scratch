package transmuter.framework;

public abstract class Factory {
  /** Creating product. */
  public final Product create(String owner) {
    final Product product = createProduct(owner);
    registerProduct(product);
    return product;
  }

  protected abstract Product createProduct(String owner);

  protected abstract void registerProduct(Product product);
}
