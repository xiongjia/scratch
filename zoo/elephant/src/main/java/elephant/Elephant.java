package elephant;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@EnableDiscoveryClient
@SpringBootApplication
public class Elephant {
  private static final Logger log = LoggerFactory.getLogger(Elephant.class);

  /** Elephant tests. */
  public static void main(String[] args) {
    log.debug("Elephant tests");
    SpringApplication.run(Elephant.class, args);
  }
}
