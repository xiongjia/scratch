package snow.misc;

import org.reflections.Reflections;
import org.reflections.scanners.ResourcesScanner;

import java.io.File;
import java.io.IOException;
import java.net.URISyntaxException;
import java.util.Enumeration;
import java.util.Set;
import java.util.regex.Pattern;
import java.util.zip.ZipEntry;
import java.util.zip.ZipFile;

public class JarUtilities {
  public static String getCodeLocation(Class source) throws URISyntaxException {
    return new File(source.getProtectionDomain().getCodeSource().getLocation().getPath()).getName();
  }

  public static void dump(String jarFilename) throws IOException {
    final ZipFile zipFile = new ZipFile(jarFilename);
    final Enumeration entries = zipFile.entries();
    while (entries.hasMoreElements()) {
      final ZipEntry entry = (ZipEntry)entries.nextElement();
      String message = String.format("Name = %s: Folder = %s, size = %d, compressed size = %d",
        entry.getName(), entry.isDirectory(), entry.getSize(), entry.getCompressedSize());
      System.out.println(message);
    }
  }

  public static void reflection() {
    final ResourcesScanner resourceScanner = new ResourcesScanner();
    final Reflections reflections = new Reflections(resourceScanner);
    final Set<String> files = reflections.getResources(Pattern.compile("snow-data.*"));
    System.out.println(files);
  }
}
