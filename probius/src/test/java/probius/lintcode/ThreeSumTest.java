package probius.lintcode;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.Matchers.empty;
import static org.hamcrest.Matchers.hasSize;

import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

public class ThreeSumTest {
  private static final Logger log = LoggerFactory.getLogger(ThreeSumTest.class);
  private final ThreeSum inst = new ThreeSum();

  @Test
  public void test() {
    assertThat(this.inst.threeSum(new int[] {}), empty());
    assertThat(this.inst.threeSum(new int[] {1, 2 }), empty());
    assertThat(this.inst.threeSum(new int[] {1, 2, 3, 4 }), empty());

    final List<List<Integer>> res = this.inst.threeSum(new int[] {-1, 0, 1, 2, -1, -4});
    log.debug("res: {}", res);
    assertThat(res, hasSize(2));
  }
}
