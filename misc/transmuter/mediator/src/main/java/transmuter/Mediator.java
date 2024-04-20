package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Mediator {
  private static final Logger log = LoggerFactory.getLogger(Mediator.class);

  private static class MediatorHolder {
    public static final Mediator INSTANCE = new Mediator();
  }

  public static  Mediator getInstance() {
    return MediatorHolder.INSTANCE;
  }

  public void showMessage(User user, String message) {
    log.debug("[{}] - {}", user.getName(), message);
  }
}
