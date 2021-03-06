package scratch.misc;

import lombok.AccessLevel;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NonNull;
import lombok.Setter;
import lombok.ToString;

public class LombokScratch {
  @Data
  @ToString(callSuper = false, includeFieldNames = false, exclude = "txt")
  @EqualsAndHashCode(callSuper = false, exclude = {"txt"})
  public static class LombokData {
    private String txt = this.getClass().getName();
    private int num = 0;
  }

  @Data(staticConstructor = "of")
  public static class LombokConstructor {
    private final String constructor;
    private int num = 0;
  }

  public static class LombokSetterGetter {
    @Setter @NonNull @Getter private String txt = this.getClass().getName();
    @Setter @Getter private int num = 0;
    @Setter(AccessLevel.PRIVATE) @Getter private String pval = "pval";
  }
}
