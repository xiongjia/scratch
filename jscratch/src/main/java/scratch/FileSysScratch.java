package scratch;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;

public class FileSysScratch {
  private static final Logger log = LoggerFactory.getLogger(FileSysScratch.class);

  /** file info. */
  public static void fileStat(String path) {
    log.debug("file stat: {}", path);
    final File src = new File(path);
    log.debug("absolute path: {}", src.getAbsolutePath());
    if (!src.exists()) {
      log.debug("cannot find file: {}",path);
      return;
    }
    log.debug("isDir: {}, isFile: {}", src.isDirectory(), src.isFile());
  }

  /** read dir. */
  public static void readDir(String path) {
    log.debug("read dir: {}", path);
    final File src = new File(path);
    if (!src.exists() || !src.isDirectory()) {
      log.debug("cannot find dir: {}",path);
      return;
    }

    final File [] children = src.listFiles();
    for (final File child: children) {
      log.debug("child: {}", child.getAbsolutePath());
    }
  }
}
