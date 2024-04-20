package transmuter;

import java.io.IOException;

public interface FileInputOutput {
  void readFromFile(String filename) throws IOException;

  void writeToFile(String filename) throws IOException;

  void setValue(String key, String value);

  String getValue(String key);
}
