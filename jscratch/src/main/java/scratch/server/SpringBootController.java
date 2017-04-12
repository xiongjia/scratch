package scratch.server;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SpringBootController {
  @RequestMapping("/")
  public String index() {
    return "Hello Spring Boot";
  }
}
