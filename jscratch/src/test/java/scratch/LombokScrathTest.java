package scratch;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotEquals;

import org.junit.Test;

import scratch.LombokScratch.LombokConstructor;
import scratch.LombokScratch.LombokData;
import scratch.LombokScratch.LombokSetterGetter;

public class LombokScrathTest {
  @Test
  public void lombokData() {
    LombokData ldata = new LombokData();
    assertEquals(ldata.getClass().getName(), ldata.getTxt());
    assertEquals(0, ldata.getNum());

    ldata.setTxt("New Text");
    ldata.setNum(100);
    assertEquals("New Text", ldata.getTxt());
    assertEquals(100, ldata.getNum());
  }

  @Test
  public void lombokEqualsAndHashCode() {
    LombokData ldata1 = new LombokData();
    LombokData ldata2 = new LombokData();

    ldata1.setTxt("data1");
    ldata1.setNum(0);
    ldata2.setTxt("data2");
    ldata2.setNum(0);
    assertEquals(ldata1, ldata2);
    ldata1.setNum(1);
    ldata2.setNum(2);
    assertNotEquals(ldata1, ldata2);
  }

  @Test
  public void lombokConstructor() {
    LombokConstructor ldata = LombokConstructor.of("test");
    ldata.setNum(100);
    assertEquals("test", ldata.getConstructor());
    assertEquals(100, ldata.getNum());
  }

  @Test
  public void lombokToString() {
    LombokData ldata = new LombokData();
    ldata.setNum(100);
    final String expected = String.format("LombokScratch.LombokData(%d)", ldata.getNum());
    assertEquals(expected, ldata.toString());
  }

  @Test
  public void lombokSetterGetter() {
    LombokSetterGetter ldata = new LombokSetterGetter();
    assertEquals(ldata.getClass().getName(), ldata.getTxt());
    assertEquals(0, ldata.getNum());
    assertEquals("pval", ldata.getPval());

    ldata.setTxt("New Text");
    ldata.setNum(100);
    assertEquals("New Text", ldata.getTxt());
    assertEquals(100, ldata.getNum());
  }

  @Test(expected = NullPointerException.class)
  public void lombokNonNull() {
    LombokSetterGetter ldata = new LombokSetterGetter();
    ldata.setTxt(null);
  }
}
