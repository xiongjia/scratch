package yahaha;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Yahaha {
  private static final Logger log = LoggerFactory.getLogger(Yahaha.class);

  /** Yahaha tests. */
  public static void main(String[] args) {
    log.debug("Yahaha tests");
    SpringApplication.run(Yahaha.class, args);
  }
}

