package snow;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import snow.data.CacheService;

@RestController
public class Controller {
  @Autowired
  private SnowConfiguration snowConfiguration;
  @Autowired
  private CacheService cacheService;

  @GetMapping("/config")
  public String getConfig() {
    return snowConfiguration.getMessage();
  }

  @GetMapping("/filter/another-config")
  public String getAnotherConfig() {
    return snowConfiguration.getAnotherMessage();
  }

  @GetMapping("/cache-data")
  public void getCacheData() {
    cacheService.test();
  }
}
