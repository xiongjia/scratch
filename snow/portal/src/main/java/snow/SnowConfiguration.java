package snow;

import javax.annotation.PostConstruct;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Component
@ConfigurationProperties(prefix = "snow")
public class SnowConfiguration {
  private static final Logger log = LoggerFactory.getLogger(SnowConfiguration.class);

  private String message;
  private String anotherMessage;

  @PostConstruct
  public void init() {
    log.debug("config: {}", this.toString());
  }

  public String getMessage() {
    return message;
  }

  public void setMessage(String message) {
    this.message = message;
  }

  public String getAnotherMessage() {
    return anotherMessage;
  }

  public void setAnotherMessage(String anotherMessage) {
    this.anotherMessage = anotherMessage;
  }

  @Override
  public String toString() {
    return "SnowConfiguration{"
        + "message='" + message + '\''
        + ", anotherMessage='" + anotherMessage + '\''
        + '}';
  }
}
