package transmuter;

import java.io.IOException;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);
  private static final String TESTFILE = "d:/tmp/test.prop";

  /** Adapter tests. */
  public static void main(String[] args) {
    log.debug("Adapter tests");

    final FileInputOutput file = new FileProperties();
    try {
      file.readFromFile(TESTFILE);
      file.setValue("key1", "val1");
      log.debug("key1 = {}", file.getValue("key1"));
      file.setValue("key2", "val2");
      file.writeToFile(TESTFILE);
    } catch (IOException e) {
      log.error("File IO exception. ", e);
    }
  }
}
