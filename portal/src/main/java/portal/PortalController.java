package portal;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PortalController {
  /** Index request. */
  @RequestMapping(value = "/api/test", method = RequestMethod.GET)
  public String index() {
    return "tmp";
  }
}

