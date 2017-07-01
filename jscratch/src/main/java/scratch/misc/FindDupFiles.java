package scratch.misc;

import java.io.File;
import java.io.IOException;
import java.nio.file.FileSystem;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.ArrayList;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.google.common.collect.Lists;
import lombok.Getter;
import lombok.Builder;

public class FindDupFiles {
  private static final Logger log = LoggerFactory.getLogger(FindDupFiles.class);
  private final FileSystem fs = FileSystems.getDefault();
  private final List<File> files = new ArrayList<File>();

  @Getter
  private final List<String> srcFiles;
  @Getter
  private boolean ignoreEmptyFiles;

  public FindDupFiles(List<String> srcFiles, boolean ignoreEmptyFiles) {
    this.srcFiles = srcFiles;
    this.ignoreEmptyFiles = ignoreEmptyFiles;
  }

  @Builder
  public static class FileInfo {
    @Getter
    private String fileName;
    @Getter
    private long fileSize;
    @Getter
    private String hash;
  }

  public static class FindDupFilesBuilder {
    private final List<String> src;
    private boolean ignoreEmptyFiles = true;

    public FindDupFilesBuilder(String []srcFiles) {
      this.src = Lists.newArrayList(srcFiles);
    }

    public FindDupFilesBuilder(String srcFile) {
      this.src = Lists.newArrayList(srcFile);
    }

    public FindDupFilesBuilder ignoreEmptyFiles(boolean ignore) {
      this.ignoreEmptyFiles = ignore;
      return this;
    }

    public FindDupFiles build() {
      return new FindDupFiles(this.src, this.ignoreEmptyFiles);
    }
  }

  private void scan(final String filename) {
    final Path path = fs.getPath(filename);
    final File src = path.toFile();
    if (!src.exists()) {
      return;
    }

    if (src.isFile()) {
      files.add(src);
      return;
    }

    try {
      Files.list(path).forEach((subPath) -> {
        scan(subPath.toAbsolutePath().toString());
      });
    } catch (IOException ignore) {
    }
  }

  public void run() {
    this.srcFiles.stream().forEach(file -> {
      log.debug("scan files in {}", file);
      scan(file);
    });

    log.debug("src files [{}]: {}", files.size(), files);
  }
}
