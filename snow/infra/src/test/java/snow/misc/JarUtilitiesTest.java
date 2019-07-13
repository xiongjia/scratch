package snow.misc;

import java.io.IOException;
import java.net.URISyntaxException;

import org.junit.Test;

public class JarUtilitiesTest {
  @Test
  public void testJarUtilities() throws URISyntaxException, IOException {
    final String result = JarUtilities.getCodeLocation(this.getClass());
    System.out.println(result);
    JarUtilities.dump("D:/var/dev/github/scratch/snow/build/portal/libs/portal.jar");
    // JarUtilities.reflection();
  }
}
