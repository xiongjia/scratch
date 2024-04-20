package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** State tests. */
  public static void main(String[] args) {
    log.debug("State tests");

    final Context context = new Context();
    final StartState startState = new StartState();
    startState.doAction(context);
    log.debug("state: {}", context.getState().toString());

    final StopState stopState = new StopState();
    stopState.doAction(context);
    log.debug("state: {}", context.getState().toString());
  }
}
