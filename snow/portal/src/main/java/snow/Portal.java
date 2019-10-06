package snow;

import java.util.Arrays;
import java.util.stream.StreamSupport;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.ServletComponentScan;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.context.event.EventListener;
import org.springframework.core.env.AbstractEnvironment;
import org.springframework.core.env.EnumerablePropertySource;
import org.springframework.core.env.Environment;
import org.springframework.core.env.MutablePropertySources;

@SpringBootApplication
//@ComponentScan(basePackages = {"snow"})
//@ServletComponentScan(basePackages = {"snow"})
public class Portal {
  private static final Logger log = LoggerFactory.getLogger(Portal.class);

  /** Potal. */
  public static void main(String[] args) {
    SpringApplication.run(Portal.class, args);
  }

  @EventListener
  public void applicationReady(ContextRefreshedEvent event) {
    log.debug("app ready");

    final ApplicationContext context = event.getApplicationContext();
    final Environment env = context.getEnvironment();
    log.debug("Active profiles: {}", Arrays.toString(env.getActiveProfiles()));
    final MutablePropertySources sources = ((AbstractEnvironment) env).getPropertySources();
    StreamSupport.stream(sources.spliterator(), false)
        .filter(ps -> ps instanceof EnumerablePropertySource)
        .map(ps -> ((EnumerablePropertySource) ps).getPropertyNames())
        .flatMap(Arrays::stream)
        .distinct()
        .forEach(prop -> {
          log.debug("{}: {}", prop, env.getProperty(prop));
        });
  }
}
