package slug;

import slug.general.SlugConfig;
import slug.general.SlugVersion;

public class Slug {
  private static class Instance {
    public static final SlugConfig CONFIG;

    static {
      CONFIG = SlugConfig.load();
    }
  }

  /** Slug tests. */
  public static SlugVersion getVersion() {
    return Instance.CONFIG.getVersion();
  }
}
