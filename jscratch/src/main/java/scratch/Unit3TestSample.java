package scratch;

import junit.framework.TestCase;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Unit3TestSample extends TestCase {
  private static final Logger log = LoggerFactory.getLogger(Unit3TestSample.class);

  public void testAction1() {
    log.debug("junit3 test action1");
    assertEquals("junit3", "junit3");
  }

  public void testAction2() {
    log.debug("junit3 test action2");
    assertEquals("junit3", "junit3");
  }
}
