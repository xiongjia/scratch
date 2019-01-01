package owl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class OwlController {
  private static final Logger log = LoggerFactory.getLogger(OwlController.class);

  @Autowired
  Elephant elephant;

  @RequestMapping("/")
  public String index() {
    log.debug("a testing api");
    final String result = elephant.getData();
    log.debug("result {}", result);
    return result;
  }
}
