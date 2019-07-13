package snow.misc;

import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Path;

import org.junit.Test;

public class JarUtilitiesTest {
  @Test
  public void testJarUtilities() throws URISyntaxException, IOException {
    final Path result = JarUtilities.getCodeLocation(this.getClass());
    System.out.println(result.toString());
    // JarUtilities.dump("D:/var/dev/github/scratch/snow/build/portal/libs/portal.jar");
    JarUtilities.reflection();
  }
}
