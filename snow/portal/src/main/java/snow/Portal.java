package snow;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.ServletComponentScan;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(basePackages = {"snow"})
@ServletComponentScan(basePackages = {"snow"})
public class Portal {
  private static final Logger log = LoggerFactory.getLogger(Portal.class);

  /** Portal. */
  public static void main(String[] args) {
    SpringApplication.run(Portal.class, args);
  }
}
