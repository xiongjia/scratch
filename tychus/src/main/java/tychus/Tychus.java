package tychus;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import tychus.discard.DiscardServer;
import tychus.echo.EchoServer;

public class Tychus {
  private static final Logger log = LoggerFactory.getLogger(Tychus.class);

  public static void main(String[] args) throws Exception {
    final String servType = (args.length == 0) ? "discard" : args[0];
    final int servPort = 8887;
    log.info("server type: {}, port: {}", servType, servPort);
    if (servType.equals("discard")) {
      log.debug("testing discard server");
      new DiscardServer(servPort).run();
    } else if (servType.equals("echo")) {
      log.debug("testing echo server");
      new EchoServer(servPort).run();
    } else {
      log.warn("bad server type: {}", servType);
    }
  }
}
