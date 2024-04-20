package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Memento tests. */
  public static void main(String[] args) {
    log.debug("Memento tests");
    final Gamer gamer = new Gamer(100);
    final Memento memento1 = gamer.createMemento();
    log.debug("memento: {}", memento1.getMoney());
    gamer.addMoney(100);
    final Memento memento2 = gamer.createMemento();
    log.debug("memento: {}", memento2.getMoney());

    gamer.restoreMemento(memento1);
    log.debug("restore: {}", gamer.getMoney());
  }
}
