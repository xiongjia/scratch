package scratch.misc;

import com.typesafe.config.Config;
import com.typesafe.config.ConfigFactory;

public class TypesafeConf {
  private final Config baseConf = ConfigFactory.load("typesafe-conf-base.conf");
  private final Config conf = ConfigFactory.load("typesafe-conf.conf").withFallback(baseConf);

  public String getConfName() {
    return conf.getString("jscratch.conf-name");
  }

  public String getStrVal(String path) {
    return conf.getString(path);
  }
}
