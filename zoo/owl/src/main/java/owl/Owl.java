package owl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
// @EnableFeignClients
public class Owl {
  private static final Logger log = LoggerFactory.getLogger(Owl.class);

  /** Elephant tests. */
  public static void main(String[] args) {
    log.debug("Owl tests");
    SpringApplication.run(Owl.class, args);
  }
}
