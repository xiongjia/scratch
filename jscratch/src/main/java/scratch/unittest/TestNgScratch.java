package scratch.unittest;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.testng.ITestContext;
import org.testng.ITestNGListener;
import org.testng.ITestResult;
import org.testng.TestNG;
import org.testng.internal.IResultListener;

public class TestNgScratch {
  private static final Logger log = LoggerFactory.getLogger(TestNgScratch.class);

  public static class TestListener implements IResultListener {
    @Override
    public void onTestStart(ITestResult result) {
      log.debug("onTestStart");
    }

    @Override
    public void onTestSuccess(ITestResult result) {
      log.debug("onTestSuccess");
    }

    @Override
    public void onTestFailure(ITestResult result) {
      log.debug("onTestFailure");
    }

    @Override
    public void onTestSkipped(ITestResult result) {
      log.debug("onTestSkipped");
    }

    @Override
    public void onTestFailedButWithinSuccessPercentage(ITestResult result) {
      log.debug("onTestFailedButWithinSuccessPercentage");
    }

    @Override
    public void onStart(ITestContext context) {
      log.debug("onStart");
    }

    @Override
    public void onFinish(ITestContext context) {
      log.debug("onFinish");
    }

    @Override
    public void onConfigurationSuccess(ITestResult itr) {
      log.debug("onConfigurationSuccess");
    }

    @Override
    public void onConfigurationFailure(ITestResult itr) {
      log.debug("onConfigurationFailure");
    }

    @Override
    public void onConfigurationSkip(ITestResult itr) {
      log.debug("onConfigurationSkip");
    }
  }

  /** test NG Runner. */
  public static void test() {
    log.debug("TestNgScratch testing");

    /* creating test NG runner */
    final TestNG testNg = new TestNG();
    testNg.setUseDefaultListeners(false);
    testNg.setVerbose(0);
    testNg.addListener((ITestNGListener)new TestListener());

    /* add test class to runner */
    final Class<?>[] tsetClasses = new Class[1];
    tsetClasses[0] = TestNgSample.class;
    testNg.setTestClasses(tsetClasses);
    testNg.run();
  }
}
