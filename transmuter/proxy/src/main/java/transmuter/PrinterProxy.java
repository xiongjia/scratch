package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class PrinterProxy implements Printable {
  private static final Logger log = LoggerFactory.getLogger(PrinterProxy.class);
  private String name;
  private Printable printer;

  public PrinterProxy(String name) {
    this.name = name;
  }

  @Override
  public void setPrinterName(String name) {
    this.name = name;
    this.printer = null;
  }

  @Override
  public String getPrinterName() {
    return name;
  }

  @Override
  public void print(String message) {
    final Printable printer = getPrinter(this.name);
    printer.print(message);
  }

  private Printable getPrinter(String name) {
    if (this.printer == null) {
      log.debug("creating new printer {}", name);
      this.printer = new Printer(name);
      return this.printer;
    }
    if (this.printer.getPrinterName().equals(name)) {
      return this.printer;
    }
    log.debug("creating new printer {}", name);
    this.printer = new Printer(name);
    return this.printer;
  }
}
