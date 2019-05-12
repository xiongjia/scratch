package sheikah;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties("yahaha")
public class ExecServiceProp {
  private String message;

  public String getMessage() {
    return message;
  }

  public void setMessage(String message) {
    this.message = message;
  }
}
