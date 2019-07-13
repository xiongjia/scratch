package snow;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import snow.data.CacheService;
import snow.misc.JarUtilities;

import java.io.FileNotFoundException;
import java.net.URISyntaxException;

@RestController
public class Controller {
  @Autowired
  private SnowConfiguration snowConfiguration;
  @Autowired
  private CacheService cacheService;

  @GetMapping("/res")
  public String getRes() throws FileNotFoundException, URISyntaxException {
    return Controller.class.getProtectionDomain()
        .getCodeSource()
        .getLocation()
        .getPath().split("\\!", 2)[0];
  }

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
