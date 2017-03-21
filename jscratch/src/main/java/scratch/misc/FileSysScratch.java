package scratch.misc;

import com.google.common.jimfs.Configuration;
import com.google.common.jimfs.Jimfs;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.io.IOException;
import java.nio.file.FileSystem;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;

public class FileSysScratch {
  private static final Logger log = LoggerFactory.getLogger(FileSysScratch.class);

  /** jimfs tests. */
  public static void jimFs() throws IOException {
    final FileSystem fs = Jimfs.newFileSystem(Configuration.unix());
    final Path foo = fs.getPath("/foo");
    Files.createDirectory(foo);
    Files.list(fs.getPath("/")).forEach((path) -> {
      log.debug("item: {}", path.toAbsolutePath().toString());
    });

    final FileSystem defaultFs = FileSystems.getDefault();
    Files.list(defaultFs.getPath("/")).forEach((path) -> {
      log.debug("default fs item: {}", path.toAbsolutePath().toString());
    });
  }

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
