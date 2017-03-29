package scratch.misc;

import lombok.Data;

import java.io.Serializable;

@Data
public class BeanTarget implements Serializable {
  private static final long serialVersionUID = 5365756737853103081L;

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
