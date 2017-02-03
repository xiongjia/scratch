package scratch;

import junit.framework.JUnit4TestAdapter;
import lombok.Data;
import lombok.Getter;

import org.junit.internal.requests.FilterRequest;
import org.junit.runner.Description;
import org.junit.runner.JUnitCore;
import org.junit.runner.Result;
import org.junit.runner.manipulation.Filter;
import org.junit.runner.notification.Failure;
import org.junit.runner.notification.RunListener;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;

public class UnitTestsScratch {
  private static final Logger log = LoggerFactory.getLogger(UnitTestsScratch.class);

  @Data(staticConstructor = "of")
  public static class TestItem {
    @Getter private final String displayName;
    @Getter private final String methodName;
  }

  /** get test items from test class. */
  public ArrayList<TestItem> getTestItems(Class<?> clazz) {
    final JUnit4TestAdapter test = new JUnit4TestAdapter(clazz);
    final Description testDesc = test.getDescription();
    final ArrayList<TestItem> testItems = new ArrayList<TestItem>();
    testDesc.getChildren().stream().forEach((child) -> {
      testItems.add(TestItem.of(child.getDisplayName(), child.getMethodName()));
    });
    return testItems;
  }

  /** run test method in the test class. */
  public Result runTest(Class<?> clazz, String testMethod) {
    final JUnitCore runner = new JUnitCore();
    runner.addListener(new RunListener() {
      @Override
      public void testRunStarted(Description description) throws Exception {
        log.info("testRun started: {}", description.getDisplayName());
      }

      @Override
      public void testRunFinished(Result result) throws Exception {
        log.info("testRun finished: {}", result.toString());
      }

      @Override
      public void testStarted(Description description) throws Exception {
        log.info("test started: {}", description.getDisplayName());
      }

      @Override
      public void testFinished(Description description) throws Exception {
        log.info("test finished: {}", description.getDisplayName());
      }

      @Override
      public void testFailure(Failure failure) throws Exception {
        log.info("test failure: {}", failure.getMessage());
      }
    });
    return runner.run(FilterRequest.aClass(clazz).filterWith(new Filter() {
      @Override
      public boolean shouldRun(Description description) {
        return testMethod == null || testMethod.equals(description.getMethodName());
      }

      @Override
      public String describe() {
        return String.format("Test %s", testMethod);
      }
    }).getRunner());
  }
}
