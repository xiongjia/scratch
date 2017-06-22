package orthrus;

import lombok.Builder;
import lombok.Data;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class OrthrusController {
  @Builder
  @Data
  public static class TestResponseEnt {
    private String data = "data";
    private int num = 100;
  }

  /** Index request. */
  @RequestMapping(value = "/api/test", method = RequestMethod.GET)
  public ResponseEntity<TestResponseEnt> index() {
    return new ResponseEntity<TestResponseEnt>(new TestResponseEnt.TestResponseEntBuilder()
        .data("SprintTest").num(666).build(),
        HttpStatus.OK);
  }
}
