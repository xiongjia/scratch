package sheikah;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import static org.assertj.core.api.Assertions.*;

@RunWith(SpringRunner.class)
@SpringBootTest("yahaha.message=testmessage")
public class ExecServiceTest {
  @Autowired
  private ExecService execService;

  @Test
  public void contextLoads() {
    System.out.println(execService.getMessage());
    assertThat(execService.getMessage()).isNotNull();
  }

  @SpringBootApplication
  static class TestConfiguration {
  }
}
