package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SingletonStaticBlock {
  private static final Logger log = LoggerFactory.getLogger(SingletonStaticBlock.class);
  private static final SingletonStaticBlock INSTANCE;

  static {
    INSTANCE = new SingletonStaticBlock();
  }

  public SingletonStaticBlock() {
    log.debug("SingletonStaticBlock construct()");
  }

  public static SingletonStaticBlock getInstance() {
    log.debug("SingletonStaticBlock getInstance()");
    return INSTANCE;
  }

  public void print() {
    System.out.println(SingletonStaticBlock.class);
  }
}
