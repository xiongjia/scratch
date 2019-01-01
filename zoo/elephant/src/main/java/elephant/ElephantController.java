package elephant;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ElephantController {
  private static final Logger log = LoggerFactory.getLogger(ElephantController.class);

  @GetMapping("/")
  public String index() {
    log.debug("a testing api");
    return "elephant test";
  }
}
