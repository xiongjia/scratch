package tatooine.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("local")
public class TestController {
  @GetMapping
  public String test() {
    return "test1024";
  }
}
