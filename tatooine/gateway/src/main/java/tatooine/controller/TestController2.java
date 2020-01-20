package tatooine.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("local2")
public class TestController2 {
  @GetMapping
  public String test() {
    return "test2048";
  }
}
