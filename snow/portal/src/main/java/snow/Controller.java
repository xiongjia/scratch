package snow;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Controller {
  @Autowired
  private SnowConfiguration snowConfiguration;

  @GetMapping("/config")
  public String getConfig() {
    return snowConfiguration.getMessage();
  }

  @GetMapping("/another-config")
  public String getAnotherConfig() {
    return snowConfiguration.getAnotherMessage();
  }
}
