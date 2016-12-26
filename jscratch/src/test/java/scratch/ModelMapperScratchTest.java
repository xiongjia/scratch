package scratch;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

import scratch.ModelMapperScratch.Source;
import scratch.ModelMapperScratch.Target;

public class ModelMapperScratchTest {
  @Test
  public void baseTest() {
    Source src = new Source();
    src.setSrcStr(this.getClass().getName());
    src.setSrcNum(100);
    Target target = ModelMapperScratch.toTarget(src);

    assertEquals(src.getSrcStr(), target.getTargetStr());
    assertEquals(src.getSrcNum(), target.getTargetNum());
  }
}
