package snow.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import snow.debugger.DumpProperties;

@RestController
@RequestMapping("api/v1/debug")
public class Debugger {
  @Autowired
  private DumpProperties dumpProperties;

  @GetMapping(value = "properties", produces = "text/plain")
  public String dumpProperties() {
    return "All Properties: \n" + dumpProperties.dumpAll();
  }
}
