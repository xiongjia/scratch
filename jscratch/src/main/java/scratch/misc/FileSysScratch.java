package scratch.misc;

import com.google.common.jimfs.Configuration;
import com.google.common.jimfs.Jimfs;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.channels.SeekableByteChannel;
import java.nio.charset.Charset;
import java.nio.file.FileSystem;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Formatter;

public class FileSysScratch {
  private static final Logger log = LoggerFactory.getLogger(FileSysScratch.class);

  private static String byteArray2Hex(final byte[] hash) {
    Formatter fmt = new Formatter();
    for (byte b : hash) {
      fmt.format("%02x", b);
    }
    final String result = fmt.toString();
    fmt.close();
    return result;
  }

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

  /** file content. */
  public static void fileData(String filename) throws IOException, NoSuchAlgorithmException {
    final FileSystem fs = FileSystems.getDefault();
    final Path path = fs.getPath(filename);
    final File file = path.toFile();
    if (!file.exists() || file.isDirectory()) {
      log.debug("cannot find file {}", filename);
      return;
    }

    if (!file.canRead()) {
      log.debug("cannot read file {}", filename);
    }


    final ByteBuffer buf = ByteBuffer.allocate(1024 * 4);
    final Charset charset = Charset.forName("US-ASCII");
    final MessageDigest md = MessageDigest.getInstance("SHA");
    final SeekableByteChannel byteChannel = Files.newByteChannel(path);
    while (byteChannel.read(buf) > 0) {
      md.update(buf);
      buf.rewind();
      log.debug("buf: {}", charset.decode(buf));
      buf.flip();
    }
    byteChannel.close();
    log.debug("sha1: {}", byteArray2Hex(md.digest()));
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
    log.debug("file size: {}", src.length());
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
