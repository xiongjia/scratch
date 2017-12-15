package scratch;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

import org.junit.Test;

import org.junit.runner.Result;
import scratch.unittest.UnitTestsScratch;
import scratch.unittest.UnitTestsScratch.TestItem;

import java.util.ArrayList;

public class UnitTestsScratchTest {
  @Test
  public void listTestItems() {
    final UnitTestsScratch testsScratch = new UnitTestsScratch();
    final ArrayList<TestItem> testItems = testsScratch.getTestItems(scratch.ScratchTest.class);

    /* The test methods name in the scratch.ScratchTest.class */
    @SuppressWarnings("serial")
    final ArrayList<String> testMethods = new ArrayList<String>() {{
        add("baseTest1");
        add("baseTest2");
      }
    };
    assertEquals(testItems.size(), testMethods.size());
    testItems.stream().forEach((testItem) -> {
      assertTrue(testMethods.contains(testItem.getMethodName()));
    });
  }

  @Test
  public void runTestItem() {
    final UnitTestsScratch testsScratch = new UnitTestsScratch();
    final Result baseTest1 = testsScratch.runTest(scratch.ScratchTest.class, "baseTest1");
    assertEquals(baseTest1.getRunCount(), 1);

    final Result allTests = testsScratch.runTest(scratch.ScratchTest.class, null);
    assertEquals(allTests.getRunCount(), 2);
  }
}
