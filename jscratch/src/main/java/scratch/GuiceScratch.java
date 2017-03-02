package scratch;

import com.google.inject.Binder;
import com.google.inject.Guice;
import com.google.inject.Injector;
import com.google.inject.Module;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Inject;

public class GuiceScratch {
  private static final Logger log = LoggerFactory.getLogger(GuiceScratch.class);

  public static class PrintService {
    public void printMessage(String message) {
      log.info("Default Print: {}", message);
    }
  }

  public static class SimplePrintService extends PrintService {
    @Override
    public void printMessage(String message) {
      log.info("Simple Print: {}", message);
    }
  }

  public static class ComplexPrintService extends PrintService {
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

  public static class TestApp {
    @Inject
    private PrintService printSvc = new PrintService();

    public void print() {
      printSvc.printMessage("Test APP");
    }
  }

  /** guice tests. */
  public static void guiceTest(boolean simplePrint) {
    final Injector injector = Guice.createInjector(new PrintServiceModule(simplePrint));
    final PrintService printSvc = injector.getInstance(PrintService.class);
    printSvc.printMessage("test");

    final TestApp testApp = injector.getInstance(TestApp.class);
    testApp.print();
  }
}