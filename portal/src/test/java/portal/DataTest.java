package portal;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

// import javax.annotation.Resource;
import java.util.List;

// @RunWith(SpringRunner.class)
// @SpringBootTest(classes = Portal.class)
public class DataTest {
  private static final Logger logger = LoggerFactory.getLogger(DataTest.class);

  // @Resource
  // private DataService dataService;


  @Test
  public void simpleTest() {
    // final List<String> data = dataService.getAllData();
    // logger.debug("data: {}", data);
    logger.debug("data: {}", "disable tests");
  }
}
