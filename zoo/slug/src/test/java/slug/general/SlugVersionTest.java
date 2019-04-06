package slug.general;

import org.junit.Test;

import static org.hamcrest.text.MatchesPattern.matchesPattern;
import static org.junit.Assert.*;

public class SlugVersionTest {
  @Test
  public void getFullVersion() {
    final SlugVersion version = new SlugVersion();
    version.setMajor(22);
    version.setMinor(33);
    assertThat(version.getFullVersion(), matchesPattern("^22.*"));
    assertThat(version.getFullVersion(), matchesPattern(".*33$"));
  }
}
