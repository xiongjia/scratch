package scratch;

import static org.junit.Assert.*;
import org.junit.Test;
import scratch.ModelMapperScratch.*;

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
