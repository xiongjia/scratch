package scratch;

import static org.junit.Assert.assertEquals;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Ignore;
import org.junit.Test;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ScratchTest {
  private static final Logger log = LoggerFactory.getLogger(ScratchTest.class);

  @BeforeClass
  public static void setupClass() {
    log.info("setupClass()");
  }

  @AfterClass
  public static void tearDownClass() {
    log.info("tearDownClass()");
  }

  @Before
  public final void setUp() {
    log.info("setUp()");
  }

  @After
  public final void tearDown() {
    log.info("tearDown()");
  }

  @Test
  public void baseTest1() {
    log.info("baseTest1()");
    assertEquals("baseTest1", "baseTest1");
  }

  @Test
  public void baseTest2() {
    log.info("baseTest2()");
  }

  @Ignore("Ignoe baseTest3")
  @Test
  public void baseTest3() {
    log.info("baseTest3()");
  }
}
