package portal;

import io.swagger.annotations.ApiOperation;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PortalController {
  /** Index request. */
  @ApiOperation("scratch")
  @RequestMapping(value = "/api/test", method = RequestMethod.GET)
  public String index() {
    return "tmp";
  }
}

