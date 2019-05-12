package sheikah;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@EnableConfigurationProperties(ExecServiceProp.class)
public class ExecService {
  private final ExecServiceProp execServiceProp;

  public ExecService(ExecServiceProp serviceProperties) {
    this.execServiceProp = serviceProperties;
  }

  public String getMessage() {
    return execServiceProp.getMessage();
  }
}
