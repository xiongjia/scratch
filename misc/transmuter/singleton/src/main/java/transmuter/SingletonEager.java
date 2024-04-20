package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SingletonEager {
  private static final Logger log = LoggerFactory.getLogger(SingletonEager.class);
  private static final SingletonEager INSTANCE = new SingletonEager();

  public SingletonEager() {
    log.debug("SingletonEager constructor()");
  }

  public static SingletonEager getInstance() {
    log.debug("SingletonEager getInstance()");
    return INSTANCE;
  }

  public void print() {
    System.out.println(SingletonEager.class);
  }
}
