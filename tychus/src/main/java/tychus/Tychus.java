package tychus;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import tychus.discard.DiscardServer;

public class Tychus {
  private static final Logger log = LoggerFactory.getLogger(Tychus.class);

  public static void main(String[] args) throws Exception {
    log.debug("testing discard server");
    new DiscardServer(8887).run();
  }
}
