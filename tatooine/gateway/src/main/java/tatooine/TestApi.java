package tatooine;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestApi {
  @GetMapping("/api/v1/test")
  public String test() {
    return "test101";
  }

  @GetMapping("/test1")
  public String test2() {
    return "test102";
  }
}
