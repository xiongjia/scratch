package scratch.misc;

import lombok.Data;

@Data
public class BeanSource {
  private String srcStr = "source";
  private int srcNum = 0;

  public BeanSource() {
    /* NOOP */
  }

  public BeanSource(String str, int num) {
    this.srcStr = str;
    this.srcNum = num;
  }
}
