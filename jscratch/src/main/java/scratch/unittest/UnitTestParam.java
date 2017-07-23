package scratch.unittest;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameter;
import org.junit.runners.Parameterized.Parameters;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;
import java.util.Collection;

@RunWith(Parameterized.class)
public class UnitTestParam {
  private static final Logger log = LoggerFactory.getLogger(UnitTestParam.class);

  @Parameter(0)
  public int first;
  @Parameter(1)
  public int second;

  /** test parameters. */
  @Parameters(name = "{index}: data({0})={1}")
  public static Collection<Object[]> data() {
    return Arrays.asList(new Object[][] {
      { 0, 0 }, { 1, 1 }, { 2, 2 }
    });
  }

  @Test
  public void testSample() {
    log.debug("test first = {}; second = {}", first, second);
  }
}
