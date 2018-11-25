package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Proxy tests. */
  public static void main(String[] args) {
    log.debug("Proxy tests");
    final Printable printer = new PrinterProxy("test1");
    printer.print("message1");
    printer.print("message2");
    printer.setPrinterName("test2");
    printer.print("message3");
  }
}
