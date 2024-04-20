package transmuter;

import java.io.Serializable;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SingletonSerializable implements Serializable {
  private static final Logger log = LoggerFactory.getLogger(SingletonSerializable.class);
  private static final long serialVersionUID = 4010632389650850251L;

  private int testValue = 10;

  private static class Holder {
    public static final SingletonSerializable INSTANCE = new SingletonSerializable();
  }

  public static SingletonSerializable getInstance() {
    log.debug("SingletonSerializable getInstance()");
    return Holder.INSTANCE;
  }

  public SingletonSerializable() {
    log.debug("SingletonSerializable construct()");
  }

  public void print() {
    System.out.print(SingletonSerializable.class);
  }

  public int getTestValue() {
    return testValue;
  }

  public void setTestValue(int testValue) {
    this.testValue = testValue;
  }

  protected Object readResolve() {
    log.debug("SingletonSerializable readResolve()");
    return getInstance();
  }
}
