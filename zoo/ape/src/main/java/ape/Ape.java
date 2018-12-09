package ape;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Ape {
  private static final Logger log = LoggerFactory.getLogger(Ape.class);

  /** Ape tests. */
  public static void main(String[] args) {
    log.debug("Ape tests");
    SpringApplication.run(Ape.class, args);
  }
}
