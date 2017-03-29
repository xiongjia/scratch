package scratch.misc;

import lombok.Data;

import java.io.Serializable;

@Data
public class BeanSource implements Serializable {
  private static final long serialVersionUID = -8242134062070572818L;

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
