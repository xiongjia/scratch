package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SingletonLazy {
  private static final Logger log = LoggerFactory.getLogger(SingletonLazy.class);
  private static volatile SingletonLazy instance = null;

  public SingletonLazy() {
    log.debug("SingletonLazy construct()");
  }

  /** Singleton Lazy - get instance. */
  public static synchronized SingletonLazy getInstance() {
    log.debug("SingletonLazy getInstance()");
    if (instance == null) {
      instance = new SingletonLazy();
    }
    return instance;
  }

  public void print() {
    System.out.println(SingletonLazy.class);
  }
}
