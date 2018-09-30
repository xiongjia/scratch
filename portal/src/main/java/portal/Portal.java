package portal;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.util.Arrays;

@SpringBootApplication
public class Portal implements CommandLineRunner {
  private static final Logger logger = LoggerFactory.getLogger(Portal.class);

  public static void main(String[] args) {
    SpringApplication.run(Portal.class, args);
  }

  @Override
  public void run(String... args)  {
    logger.debug("Portal runner: {}", Arrays.toString(args));
  }
}
