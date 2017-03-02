package scratch;

import com.google.inject.Binder;
import com.google.inject.Guice;
import com.google.inject.Injector;
import com.google.inject.Module;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class GuiceScratch {
  private static final Logger log = LoggerFactory.getLogger(GuiceScratch.class);

  public interface PrintService {
    public void printMessage(String message);
  }

  public static class SimplePrintService implements PrintService {
    @Override
    public void printMessage(String message) {
      log.info("Simple Print: {}", message);
    }
  }

  public static class ComplexPrintService implements PrintService {
    @Override
    public void printMessage(String message) {
      log.info("Complex Print: {}", message);
    }
  }

  public static class PrintServiceModule implements Module {
    final Class<? extends PrintService> printSvcClass;

    public PrintServiceModule(boolean simplePrint) {
      this.printSvcClass = simplePrint ? SimplePrintService.class : ComplexPrintService.class;
    }

    @Override
    public void configure(Binder binder) {
      binder.bind(PrintService.class).to(printSvcClass);
    }
  }

  /** guice tests. */
  public static void guiceTest(boolean simplePrint) {
    final Injector injector = Guice.createInjector(new PrintServiceModule(simplePrint));
    final PrintService printSvc = injector.getInstance(PrintService.class);
    printSvc.printMessage("test");
  }
}