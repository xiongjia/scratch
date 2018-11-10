package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SingletonInnerHolder {
  private static final Logger log = LoggerFactory.getLogger(SingletonInnerHolder.class);

  private static class Holder {
    public static final SingletonInnerHolder INSTANCE = new SingletonInnerHolder();
  }

  public SingletonInnerHolder() {
    log.debug("SingletonInnerHolder constructor()");
  }

  public static SingletonInnerHolder getInstance() {
    log.debug("SingletonInnerHolder getInstance()");
    return Holder.INSTANCE;
  }

  public void print() {
    System.out.println(SingletonInnerHolder.class);
  }
}
