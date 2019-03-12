package owl;

import io.swagger.annotations.ApiOperation;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import javax.validation.constraints.Size;


@Controller
public class OwlController {
  private static final Logger log = LoggerFactory.getLogger(OwlController.class);

  // @Autowired
  // Elephant elephant;

  @ApiOperation(value = "scratch", notes = "notes")
  @RequestMapping(value = "/test", method = RequestMethod.GET)
  @ResponseStatus(code = HttpStatus.BAD_REQUEST)
  public String index(@RequestParam(required = false) @Size(min = 1, max = 5) String name) {
//    if (name == null) {
//      throw new IllegalArgumentException("The 'name' parameter must not be null or empty");
//    }
    return "home";
    // log.debug("a testing api");
    // final String result = elephant.getData();
    // log.debug("result {}", result);
    // return result;

    // return "test";
  }

  @RequestMapping(value = "/login", method = RequestMethod.GET)
  public String login() {
    return "login";
  }
}
