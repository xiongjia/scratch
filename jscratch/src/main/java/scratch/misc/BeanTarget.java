package scratch.misc;

import lombok.Data;

@Data
public class BeanTarget {
  private String targetStr = "target";
  private int targetNum = 0;

  public BeanTarget() {
    /* NOOP */
  }

  public BeanTarget(String str, int num) {
    this.targetStr = str;
    this.targetNum = num;
  }
}
