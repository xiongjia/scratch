package sheikah;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;

@Configuration
@ComponentScan(basePackages = "sheikah")
public class ModuleConfig {
  private static final Logger log = LoggerFactory.getLogger(ModuleConfig.class);

  @PostConstruct
  public void postConstruct() {
    log.info("MODULE LOADED!");
  }
}
