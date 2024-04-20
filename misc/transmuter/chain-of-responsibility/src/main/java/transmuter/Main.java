package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Chain of responsibility tests. */
  public static void main(String[] args) {
    log.debug("Chain of responsibility tests");

    final Support alice = new NoSupport("Alice");
    final Support bob = new LimitSupport("Bob", 100);
    final Support charlie = new SpecialSupport("Charlie", 429);
    final Support diana = new LimitSupport("Diana", 200);
    final Support elmo = new OddSupport("Elmo");
    final Support fred = new LimitSupport("Fred", 300);
    alice.setNext(bob).setNext(charlie).setNext(diana).setNext(elmo).setNext(fred);
    for (int i = 0; i < 500; i += 33) {
      alice.support(new Trouble(i));
    }
  }
}
