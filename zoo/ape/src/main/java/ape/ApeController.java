package ape;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ApeController {
  private static final Logger log = LoggerFactory.getLogger(ApeController.class);

  @RequestMapping("/")
  public String index() {
    log.debug("a testing api");
    return "hello";
  }
}
