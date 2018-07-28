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
  public static class TestData {
    @Builder.Default
    private String data = "testdata";
    @Builder.Default
    private int num = 100;
  }

  /** Index request. */
  @RequestMapping(value = "/api/testdata", method = RequestMethod.GET)
  public ResponseEntity<TestData> index() {
    return new ResponseEntity<TestData>(new TestData.TestDataBuilder()
        .data("SprintTest").num(666).build(),
        HttpStatus.OK);
  }
}
