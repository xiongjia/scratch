package slug.general;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;

import java.io.IOException;
import java.io.InputStream;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SlugConfig {
  private static final Logger log = LoggerFactory.getLogger(SlugConfig.class);
  private SlugVersion version;

  public SlugVersion getVersion() {
    return version;
  }

  public void setVersion(SlugVersion version) {
    this.version = version;
  }

  private static InputStream loadResource(String name) {
    return SlugConfig.class.getClassLoader().getResourceAsStream(name);
  }

  private static SlugConfig loadDefault() {
    final SlugConfig defaultConf = new SlugConfig();
    final SlugVersion version = new SlugVersion();
    version.setMajor(1);
    version.setMinor(0);
    defaultConf.setVersion(version);
    return defaultConf;
  }

  /** load slug config. */
  public static SlugConfig load() {
    log.debug("loading slug config");
    try (final InputStream content = loadResource("slug.yaml")) {
      final ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
      final SlugConfig config = mapper.readValue(content, SlugConfig.class);
      return config;
    } catch (IOException err) {
      log.error("load 'slug.yaml' error", err);
      return loadDefault();
    }
  }
}
