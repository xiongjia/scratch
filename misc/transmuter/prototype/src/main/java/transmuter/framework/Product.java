package transmuter.framework;

public abstract class Product implements Cloneable {
  public abstract void use(String str);

  /** Create clone object. */
  public Product createClone() {
    try {
      return (Product)clone();
    } catch (CloneNotSupportedException e) {
      e.printStackTrace();
      return null;
    }
  }
}
