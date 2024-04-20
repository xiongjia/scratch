package transmuter.pagemaker;

import java.util.HashMap;
import java.util.Map;

public class Database {
  private static final Map<String, String> users;

  static {
    users = new HashMap<>();
    users.put("test1@test.com", "test1");
    users.put("test2@test.com", "test2");
  }

  private Database() {
    // NOP
  }

  public static String getUsername(String mailAddr) {
    return users.get(mailAddr);
  }
}
