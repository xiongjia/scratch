package owl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import slug.CommandExecution;
import slug.Slug;

@SpringBootApplication
// @EnableFeignClients
public class Owl {
  private static final Logger log = LoggerFactory.getLogger(Owl.class);

  /** Elephant tests. */
  public static void main(String[] args) {
    final Slug slug = new Slug();
    final String version = slug.getVersion();
    slug.listTag();

    final CommandExecution cmdExec = new CommandExecution();
    final String execResult = cmdExec.runTest();


    log.debug("Owl tests");
    SpringApplication.run(Owl.class, args);
  }
}
