package scratch;

import org.junit.Test;
import static org.junit.Assert.*;

public class LombokScrathTest {
    @Test
    public void lombokData() {
    	LombokScratch.LombokData ldata = new LombokScratch.LombokData();
        assertEquals(ldata.getClass().getName(), ldata.getTxt());
        assertEquals(0, ldata.getNum());
        
        ldata.setTxt("New Text");
        ldata.setNum(100);
        assertEquals("New Text", ldata.getTxt());
        assertEquals(100, ldata.getNum());
    }

    @Test
    public void lombokSetterGetter() {
    	LombokScratch.LombokSetterGetter ldata = new LombokScratch.LombokSetterGetter();
        assertEquals(ldata.getClass().getName(), ldata.getTxt());
        assertEquals(0, ldata.getNum());
        
        ldata.setTxt("New Text");
        ldata.setNum(100);
        assertEquals("New Text", ldata.getTxt());
        assertEquals(100, ldata.getNum());   	
    }
}
